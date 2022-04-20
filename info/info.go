// Package info module to gather information about users, about servers, etc...
package info

import "github.com/bwmarrin/discordgo"

type ServerInfo struct {
	ID           string
	Name         string
	Description  string
	VanityURL    string
	AFKChannelID string
	AFKTimeout   string
	ApproxMember string
	ApproxPresen string
    MaxMembers   string
	MaxPresence  string


}

// UserServerInfo This function is able to list information about the user.
//Returns a struct to be able to work with it
//TODO: make it work with an appcommand
func UserServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate) ServerInfo {

	for _, guildStruct := range s.State.Guilds {
		if guildStruct.ID == i.Interaction.GuildID {
			return guildStruct.
		}
	}

	return nil
}

// UserAdminInfo This function should be able to detect the information about a user by his ID
func UserAdminInfo(s *discordgo.Session, i *discordgo.InteractionCreate, cmd []interface{}) {

}

// GetVanityServerInvite get the
func GetVanityServerInvite(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	var inviteLink string
	for _, g := range s.State.Guilds {
		if g.ID == i.Interaction.GuildID {
			inviteLink = g.VanityURLCode
		}
	}
	return inviteLink
}
