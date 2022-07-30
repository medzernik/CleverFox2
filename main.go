//This is the main setup point for the Discord bot.
//For the correct functioning, the core package group is required (main, config, logging and command packages)
//You can configure using additional modules in the config or on the fly using the bot.
//Additional modules will become available as development progresses
package main

import (
	"CleverFox2/command"
	"CleverFox2/config"
	"CleverFox2/logging"
	"CleverFox2/spinner"
	"CleverFox2/tviewsystem"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//Initialize the config
	config.LoadConfig()
	logging.Log.Traceln("Loaded the config file.")

	//Initialize the logging system
	errLogging := logging.StartLogging()
	// If the log can't be created, use stdout.
	if errLogging != nil {
		logrus.Errorln("Failed to log to file, using default stderr")
	} else {
		logging.Log.Traceln("Loaded the logging system")
	}

	// Create a new Discord session using the provided bot token. Panic if failed.
	logging.Log.Infoln("Starting the bot...")

	s, err := discordgo.New("Bot " + config.Cfg.ServerInfo.ServerToken)
	if err != nil {
		logging.Log.Panicln("error creating New session: ", err)
	}
	//Set the max message cache to 20 messages.
	s.State.MaxMessageCount = 20

	//Get the intents that are needed
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	s.Identify.Token = config.Cfg.ServerInfo.ServerToken
	s.Identify.LargeThreshold = 250

	go spinner.StartSpin(spinner.Finish, "Initializing the session to Discord...")
	//fmt.Println("\nInitializing the session to Discord...")
	err = s.Open()
	if err != nil {
		println(err)
	}
	spinner.Finish <- struct{}{}

	//Start the listening of the other functions
	go spinner.StartSpin(spinner.Finish, "Initializing commands.")
	go command.InitializeCommands(s)
	spinner.Finish <- struct{}{}

	time.Sleep(500 * time.Millisecond)
	if err := tviewsystem.StartGUI(); err != nil {
		logrus.Errorf("error starting the GUI. STDOUT only. %s", err)
		// Wait here until CTRL-C or other term signal is received.
		fmt.Println("Bot is now running in STDOUT mode only. Press CTRL+C to exit.")

		KillSignal := make(chan os.Signal, 1)
		signal.Notify(KillSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-KillSignal

		// Cleanly close down the Discord session.
		err2 := s.Close()
		if err2 != nil {
			logging.Log.Panicln("Error closing the session: ", err2)
		}
	}

}
