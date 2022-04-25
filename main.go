//This is the main setup point for the Discord bot.
//For the correct functioning, the core package group is required (main, config, logging and command packages)
//You can configure using additional modules in the config or on the fly using the bot.
//Additional modules will become available as development progresses
package main

import (
	"CleverFox2/command"
	"CleverFox2/config"
	"CleverFox2/logging"
	"CleverFox2/passive"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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
		logging.Log.Panicln("error creating Discord session: ", err)
	}

	//Get the intents that are needed
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	s.Identify.Token = config.Cfg.ServerInfo.ServerToken
	s.Identify.LargeThreshold = 250

	// Open a websocket connection to Discord and begin listening. Panic if failed.
	err = s.Open()
	if err != nil {
		logging.Log.Panicln("Error opening the websocket!: ", err)
	}

	//Start the listening of the other functions
	//responder.RegisterPlugin(s)
	command.InitializeCommands(s)
	go passive.DeleteCringe(s)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	KillSignal := make(chan os.Signal, 1)
	signal.Notify(KillSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-KillSignal

	// Cleanly close down the Discord session.
	err2 := s.Close()
	if err2 != nil {
		logging.Log.Panicln("Error closing the session: ", err2)
	}

}
