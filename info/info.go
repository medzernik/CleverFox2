package info

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

// GetCurrentServers potentially unneeded
func GetCurrentServers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, j := range s.State.Ready.Guilds {
		s.ChannelMessageSend(i.ChannelID, j.ID)
	}
}

// GetThisServer potentially unneeded
func GetThisServer(s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {
	for _, j := range s.State.Ready.Guilds {
		if i.GuildID == j.ID {
			return j.ID, nil
		}
	}
	return "", errors.New("cannot locate the server")
}

func GetServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {

	return "", errors.New("cannot locate the server")
}
