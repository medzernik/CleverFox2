// Package info module to gather information about users, about servers, etc...
package info

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type ServerInfo struct {
	ID                string
	Name              string
	Description       string
	VanityURL         string
	AFKChannelID      string
	AFKTimeout        int
	ApproxMembers     int
	ApproxPresence    int
	MaxMembers        int
	MaxPresence       int
	MFALevel          discordgo.MfaLevel
	NSFWLevel         discordgo.GuildNSFWLevel
	PeopleSubscribed  int
	NitroBoostTier    discordgo.PremiumTier
	Unavailable       bool
	Icon              string
	VerificationLevel discordgo.VerificationLevel
	Channels          []*discordgo.Channel
}
type ServerInfoString struct {
	ID                string
	Name              string
	Description       string
	VanityURL         string
	AFKChannelID      string
	AFKTimeout        string
	ApproxMembers     string
	ApproxPresence    string
	MaxMembers        string
	MaxPresence       string
	MFALevel          string
	NSFWLevel         string
	PeopleSubscribed  string
	NitroBoostTier    string
	Unavailable       string
	Icon              string
	VerificationLevel string
	Channels          []string
}

// UserServerInfo This function is able to list information about the user.
//Returns a struct to be able to work with it
//TODO: make it work with an appcommand
//TODO: making it work as a method to format the string result on demand before printing.
func UserServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate) (ServerInfo, error) {
	serverIndex, err := FindServerInArray(s, i.Interaction.GuildID)
	if err != nil {
		return ServerInfo{}, err
	}

	serverStruct := ServerInfo{
		ID:                s.State.Guilds[serverIndex].ID,
		Name:              s.State.Guilds[serverIndex].Name,
		Description:       s.State.Guilds[serverIndex].Description,
		VanityURL:         s.State.Guilds[serverIndex].VanityURLCode,
		AFKChannelID:      s.State.Guilds[serverIndex].AfkChannelID,
		AFKTimeout:        s.State.Guilds[serverIndex].AfkTimeout,
		ApproxMembers:     s.State.Guilds[serverIndex].ApproximateMemberCount,
		ApproxPresence:    s.State.Guilds[serverIndex].ApproximatePresenceCount,
		MaxMembers:        s.State.Guilds[serverIndex].MaxMembers,
		MaxPresence:       s.State.Guilds[serverIndex].MaxPresences,
		MFALevel:          s.State.Guilds[serverIndex].MfaLevel,
		NSFWLevel:         s.State.Guilds[serverIndex].NSFWLevel,
		PeopleSubscribed:  s.State.Guilds[serverIndex].PremiumSubscriptionCount,
		NitroBoostTier:    s.State.Guilds[serverIndex].PremiumTier,
		Unavailable:       s.State.Guilds[serverIndex].Unavailable,
		Icon:              s.State.Guilds[serverIndex].Icon,
		VerificationLevel: s.State.Guilds[serverIndex].VerificationLevel,
		Channels:          s.State.Guilds[serverIndex].Channels,
	}

	return serverStruct, nil
}

func (serverInfo *ServerInfo) toString(string) ServerInfoString {
	return ServerInfoString{
		ID:                "",
		Name:              "",
		Description:       "",
		VanityURL:         "",
		AFKChannelID:      "",
		AFKTimeout:        "",
		ApproxMembers:     "",
		ApproxPresence:    "",
		MaxMembers:        "",
		MaxPresence:       "",
		MFALevel:          "",
		NSFWLevel:         "",
		PeopleSubscribed:  "",
		NitroBoostTier:    "",
		Unavailable:       "",
		Icon:              "",
		VerificationLevel: "",
		Channels:          nil,
	}

}

// FindServerInArray Find a server by its ID
func FindServerInArray(s *discordgo.Session, guildID string) (int, error) {
	for guildId := range s.State.Guilds {
		if s.State.Guilds[guildId].ID == guildID {
			return guildId, nil
		}
	}
	return -1, errors.New("No server exists")
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
