// Package Info handles embed logic which is the main way for the bot to communicate with a Discord server.
// It also handles basic logic of converting the formatting of userIDs to strings and back.
package Info

import (
	"CleverFox2/config"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Define the iota for an "enum" of the message state
const (
	OK = iota
	ERROR
	WARNING
	SYNTAX
	AUTHENTICATION
	AUTOCORRECTING
)

// Define the type for the embed to process
type (
	// embedStatus struct to use for easier classification of embed types
	embedStatus struct {
		statusNumber int
		statusText   string
		statusColor  int
	}
	// EmbedInfo Embed info struct used for sending an embed
	EmbedInfo struct {
		message   string
		status    embedStatus
		title     string
		embedType discordgo.EmbedType
	}
)

// NewEmbedRich Creates a new embed using the status and message information required
func (self *EmbedInfo) NewEmbedRich(status int, message string, title ...string) *EmbedInfo {
	//case fill the status message
	titleProcessed := "OK"

	if len(title) != 0 {
		titleProcessed = ""
		for _, j := range title {
			titleProcessed = fmt.Sprintf("%s %s", strings.ToUpper(titleProcessed), j)
		}
	}

	switch status {
	case 0:
		self.status.statusNumber = status
		self.status.statusColor = 3066993
		self.status.statusText = titleProcessed

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
// TODO: Make it a variadic anonymous function that can be done using variadic argument
func (self *EmbedInfo) SendToChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get the author data from the config
	author := discordgo.MessageEmbedAuthor{
		URL:          "",
		Name:         config.Cfg.ServerInfo.BotName,
		IconURL:      config.Cfg.ServerInfo.BotLogo,
		ProxyIconURL: "",
	}

	// Create the embed based on incoming information (from the previous method call) and sending it to the channel provided.
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

	// Send a message as an embed.
	s.ChannelMessageSendEmbed(i.ChannelID, &embed)
}
