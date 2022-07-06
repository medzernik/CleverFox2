// Package Info handles embed logic which is the main way for the bot to communicate with a Discord server.
// It also handles basic logic of converting the formatting of userIDs to strings and back.
package Info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// PrintBotStatus Prints out on how many servers the bot is currently running and the author name
func PrintBotStatus(s *discordgo.Session, i *discordgo.InteractionCreate) string {

	switch len(s.State.Guilds) {
	case 1:
		return fmt.Sprintln("The bot is running on ", len(s.State.Guilds), " server")
	default:
		return fmt.Sprintln("The bot is running on ", len(s.State.Guilds), " servers")
	}
}
