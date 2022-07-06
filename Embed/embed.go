// Package Embed is part of the core group
package Embed

import (
	"CleverFox2/config"
	"github.com/bwmarrin/discordgo"
	"time"
)

//Define the iota for an "enum" of the message state
const (
	OK = iota
	ERROR
	WARNING
	SYNTAX
	AUTHENTICATION
	AUTOCORRECTING
)

//Define the type for the embed to process
type (
	//EmbedStatus struct to use for easier classification of embed types
	embedStatus struct {
		statusNumber int
		statusText   string
		statusColor  int
	}
	// EmbedInfo Embed info struct used for sending an embed
	EmbedInfo struct {
		message   string
		status    embedStatus
		embedType discordgo.EmbedType
	}
)

// NewEmbedRich Creates a new embed using the status and message information required
func (self *EmbedInfo) NewEmbedRich(status int, message string) *EmbedInfo {
	//case fill the status message
	switch status {
	case 0:
		self.status.statusNumber = status
		self.status.statusColor = 3066993
		self.status.statusText = "OK"

	case 1:
		self.status.statusNumber = status
		self.status.statusColor = 15158332
		self.status.statusText = ":bangbang: ERROR"

	case 2:
		self.status.statusNumber = status
		self.status.statusColor = 15105570
		self.status.statusText = ":warning: WARNING"

	case 3:
		self.status.statusNumber = status
		self.status.statusColor = 3447003
		self.status.statusText = ":question: SYNTAX"

	case 4:
		self.status.statusNumber = status
		self.status.statusColor = 15105570
		self.status.statusText = ":no_entry: AUTHENTICATION"

	case 5:
		self.status.statusNumber = status
		self.status.statusColor = 16776960
		self.status.statusText = ":wrench: AUTOCORRECTING"

	default:
		self.status.statusColor = 800080
		self.status.statusText = "UNDEFINED"
		self.message = "UNDEFINED"
		self.embedType = discordgo.EmbedTypeRich
		return self
	}

	//Fill out the rest of the struct
	self.embedType = discordgo.EmbedTypeRich
	self.message = message
	return self
}

// SendToChannel Sends the embed to the channel. Requires the session and interaction information
func (self *EmbedInfo) SendToChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//Fixed author message.
	author := discordgo.MessageEmbedAuthor{
		URL:          "",
		Name:         config.Cfg.ServerInfo.BotName,
		IconURL:      config.Cfg.ServerInfo.BotLogo,
		ProxyIconURL: "",
	}

	var embedArray []discordgo.MessageEmbed
	//MessageEmbed info
	//Thinking of adding timestamp time.Now().Format(time.RFC3339)
	embed := discordgo.MessageEmbed{
		URL:         "",
		Type:        self.embedType,
		Title:       self.status.statusText,
		Description: self.message,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       self.status.statusColor,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      &author,
		Fields:      nil,
	}

	embedArray = append(embedArray, embed)

	//Send a message as an embed
	s.ChannelMessageSendEmbed(i.ChannelID, &embed)

}

func ToChannelID() string {
	return ""
}
func RemoveChannelID() string {
	return ""
}
func ToMention() string {
	return ""
}
func RemoveMention() string {
	return ""
}
