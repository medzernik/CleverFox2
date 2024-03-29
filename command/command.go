package command

import (
	"CleverFox2/Info"
	"CleverFox2/logging"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// This is the setup part for the commands where the commands are defined and listed and also input sanitised by Discord.
var (
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name: "basic-command",
			// All commands and options must have a description
			// Commands/options without description will fail the registration
			// of the command.
			Description: "Basic command",
		},
		{
			Name:        "iban-to-number",
			Description: "Test command placeholder",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "IBAN",
					Required:    true,
				},
			},
		},
		{
			Name:        "number-to-iban",
			Description: "Test command placeholder",
		},
		{
			Name:                     "permission-overview",
			Description:              "Command for demonstration of default command permissions",
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
		},
		{
			Name:        "kill",
			Description: "Kills the HOSTING bot remotely!",
		},
		{
			Name:        "invite",
			Description: "Gets the vanity URL of the server, if it exists",
		},
		{
			Name:        "serverinfo",
			Description: "Gets all the available information about the server",
		},
		{
			Name:        "botstatus",
			Description: "Gets all the available information about the bot",
		},
		{
			Name:        "isadmin",
			Description: "Is user an admin?",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-mention",
					Description: "user-mention",
					Required:    true,
				},
			},
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
		"iban-to-number": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})
			var result Info.EmbedInfo

			iban, err := ParseIBAN(i.ApplicationCommandData().Options[0].StringValue())

			if err != nil {
				logging.Log.Debug("Error parsing IBAN: ", err)
				result.NewEmbedRich(Info.ERROR, fmt.Sprint(err.Error())).SendToChannel(s, i)
				return
			}

			accountNumber, err := IBANtoAccountNumber(iban)
			if err != nil {
				logging.Log.Debug("Error parsing IBAN: ", err)
				result.NewEmbedRich(Info.ERROR, fmt.Sprint(err.Error())).SendToChannel(s, i)
				return
			}

			result.NewEmbedRich(Info.OK, accountNumber.ParseToString()).SendToChannel(s, i)

		},
		"number-to-iban": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			GenerateQRCode()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Placeholder content",
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

			//setup a button that is disabled until rechecked
			var disabled bool = false

			//check whether the owner of the server is the same as the user who sent the command
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
									Emoji: &discordgo.ComponentEmoji{
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

		},
		//termination function itself - only used for terminating the bot specifically, therefore as an anonymous function
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

		},
		"serverinfo": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			guild_preview, err := s.GuildPreview(i.GuildID)
			if err != nil {
				logging.Log.Error(err)
			}

			result := Info.NewEmbed().
				SetTitle(guild_preview.Name).
				SetDescription(guild_preview.Description).
				AddField("🪪 Server ID", guild_preview.ID).
				AddField("🧑‍🤝‍🧑 Approx. members", fmt.Sprint((guild_preview.ApproximateMemberCount))).
				AddField("🔧 Server features", fmt.Sprintln(guild_preview.Features)).
				SetColor(3066993).
				SetThumbnail(guild_preview.IconURL("32")).
				SetFooter(time.Now().Format(time.RFC1123)).
				InlineAllFields()

			s.ChannelMessageSendEmbed(i.ChannelID, result.MessageEmbed)

		},
		"botstatus": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			var result Info.EmbedInfo
			result.NewEmbedRich(Info.OK, Info.PrintBotStatus(s, i)).SendToChannel(s, i)

		},
		"isadmin": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})

			var result Info.EmbedInfo

			var User Info.UserID = Info.UserID(i.ApplicationCommandData().Options[0].UserValue(s).ID)
			isAdmin, err := IsAdmin(s, i, User)
			if err != nil {
				result.NewEmbedRich(Info.ERROR, "Error getting user: "+fmt.Sprintf(err.Error())).SendToChannel(s, i)
				return
			}
			if isAdmin {
				resultText := "User " + User.ToUserMention().ToString() + " is an admin"
				result.NewEmbedRich(Info.OK, resultText, "true").SendToChannel(s, i)
			} else {
				resultText := "User " + User.ToUserMention().ToString() + " is not an admin"
				result.NewEmbedRich(Info.OK, resultText, "false").SendToChannel(s, i)
			}

			fmt.Println("Is admin:", isAdmin)

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

// Start registers commands to be created
func Start(s *discordgo.Session) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		for _, guildID := range s.State.Guilds {
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildID.ID, v)
			if err != nil {
				logging.Log.Warn("Cannot create:", v.Name, "command:", v.Name, err)
				// fmt.Printf("Cannot create '%v' command: %v", v.Name, err)
			}
			registeredCommands[i] = cmd
			logging.Log.Trace("Registered command:", v.Name)
		}

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
