// This is the main setup point for the Discord bot.
// For the correct functioning, the core package group is required (main, config, logging and command packages)
// You can configure using additional modules in the config or on the fly using the bot.
// Additional modules will become available as development progresses
package main

import (
	"CleverFox2/command"
	"CleverFox2/config"
	"CleverFox2/logging"
	"CleverFox2/tviewsystem"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {

	go tviewsystem.StartGUI()
	//Initialize the config
	config.LoadConfig()
	logging.Log.Traceln("Loaded the config file.")
	tviewsystem.StatusPush("Loaded the config file.")

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
	tviewsystem.StatusPush("Starting the bot...")

	discord_session, err := discordgo.New("Bot " + config.Cfg.ServerInfo.ServerToken)
	if err != nil {
		logging.Log.Panicln("error creating New session: ", err)
	}
	//Set the max message cache to 20 messages.
	discord_session.State.MaxMessageCount = 20

	//Get the intents that are needed
	discord_session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	discord_session.Identify.Token = config.Cfg.ServerInfo.ServerToken
	discord_session.Identify.LargeThreshold = 250

	discord_session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageCreate) {
		var guild_name string

		for _, guild := range s.State.Guilds {
			if guild.ID == e.GuildID {
				guild_name = guild.Name
			}
		}

		tviewsystem.MainViewPush(fmt.Sprintf("[%s] %s %s %s", guild_name, e.Author, e.Timestamp.Format(time.Kitchen), e.Content))
	})

	discord_session.AddHandler(func(s *discordgo.Session, e *discordgo.MessageUpdate) {
		var guild_name string

		for _, guild := range s.State.Guilds {
			if guild.ID == e.GuildID {
				guild_name = guild.Name
			}
		}

		tviewsystem.MainViewPush(fmt.Sprintf("[EDITED]:[%s] %s %s %s", guild_name, e.Author, e.Timestamp.Format(time.Kitchen), e.Content))
	})

	//go spinner.StartSpin(spinner.Finish, "Initializing the session to Discord...")
	tviewsystem.StatusPush("Initializing the session to Discord...")
	//fmt.Println("\nInitializing the session to Discord...")
	err = discord_session.Open()
	if err != nil {
		println(err)
	}
	//spinner.Finish <- struct{}{}

	//Start the listening of the other functions
	//go spinner.StartSpin(spinner.Finish, "Initializing commands.")
	tviewsystem.StatusPush("Initializing commands.")
	command.InitializeCommands(discord_session)
	//spinner.Finish <- struct{}{}

	//time.Sleep(500 * time.Millisecond)

	tviewsystem.StatusPush("Bot is now running. Press ESC and confirm to quit.")

	//Initialize the GUI system. If this fails, fallback to the stdout legacy printouts.

	// if err := tviewsystem.StartGUI(); err != nil {
	// 	logrus.Errorf("error starting the GUI. STDOUT only. %s", err)
	// 	// Wait here until CTRL-C or other term signal is received.
	// 	fmt.Println("Bot is now running in STDOUT mode only. Press CTRL+C to exit.")

	// 	KillSignal := make(chan os.Signal, 1)
	// 	signal.Notify(KillSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// 	<-KillSignal

	// 	// Cleanly close down the Discord session.
	// 	err2 := discord_session.Close()
	// 	if err2 != nil {
	// 		logging.Log.Panicln("Error closing the session: ", err2)
	// 	}
	// }

	go GetAndPostMessage()

	KillSignal := make(chan os.Signal, 1)
	signal.Notify(KillSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-KillSignal

}

func GetAndPostMessage() {

}
