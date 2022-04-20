package command

import (
	"CleverFox2/info"
	"CleverFox2/logging"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

//This is the setup part for the commands where the commands are defined and listed and also input sanitised by Discord.
var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
		{
			Name:        "kill",
			Description: "Kills the HOSTING bot remotely!",
		},
	}

	//This part of the command process actually lists the logic and responses of the commands. The "name" must match the "name" of the above section.
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
		"kill": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			guildInfo, err := s.Guild(i.Interaction.GuildID)
			if err != nil {
				logging.Log.Error("Error getting guild info", err)
				return
			}

			var disabled bool = true

			if guildInfo.OwnerID == i.Member.User.ID {
				logging.Log.Info("Bot shutting down at the request of the owner...")

				disabled = false

			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Press the button to kill the bot. Works only for the owner.",
					Flags:   1 << 6,
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Emoji: discordgo.ComponentEmoji{
										Name: "⚠️",
									},
									Label:    "Kill the bot",
									Style:    discordgo.DangerButton,
									CustomID: "terminate",
									Disabled: disabled,
								},
							},
						},
					},
				},
			})
			if err != nil {
				logging.Log.Warn(err)
				fmt.Println(err.Error())
				return
			}

			return
		},

		"terminate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			logging.Log.Info("Terminating session.")

			time.Sleep(1 * time.Second)
			//Kill the bot
			err2 := s.Close()
			if err2 != nil {
				logging.Log.Panicln("Error closing the session: ", err2)
			}
			time.Sleep(2 * time.Second)
			os.Exit(0)
		},
		//get the invite for the server, if it has a vanity URL, print that
		"invite": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			info.GetServerInvite(s, i)
			return
		},
	}
)

// InitializeCommands function runs to initialize commands in the given session.
func InitializeCommands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:

			if h, ok := commandHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
	Start(s)
}

//Start registers commands to be created
func Start(s *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, s.State.Guilds[0].ID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	/*
		if *RemoveCommands {
			log.Println("Removing commands...")
			// // We need to fetch the commands, since deleting requires the command ID.
			// // We are doing this from the returned commands on line 375, because using
			// // this will delete all the commands, which might not be desirable, so we
			// // are deleting only the commands that we added.
			// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
			// if err != nil {
			// 	log.Fatalf("Could not fetch registered commands: %v", err)
			// }

			for _, v := range registeredCommands {
				err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
				if err != nil {
					log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
				}
			}
		}

	*/
}
