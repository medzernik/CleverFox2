// Package Info handles embed logic which is the main way for the bot to communicate with a Discord server.
// It also handles basic logic of converting the formatting of userIDs to strings and back.

package Info

import (
	"strings"
)

// UserID User string type with methods to edit it in place.
type UserID string

// UserIDConversion No use for this yet.
type UserIDConversion interface {
	ToString()
	ToUserMention()
	ToRoleMention()
	ToChannelMention()
	strings.Replacer
}

// ToID Strips the mention to the bare string ID that most functions need
func (self *UserID) ToID() *UserID {
	internalMutation := string(*self)

	internalMutation = strings.Replace(internalMutation, "<", "", 1)
	internalMutation = strings.Replace(internalMutation, ">", "", 1)
	internalMutation = strings.Replace(internalMutation, "!", "", 1)
	internalMutation = strings.Replace(internalMutation, "@", "", 1)
	internalMutation = strings.Replace(internalMutation, "&", "", 1)
	internalMutation = strings.Replace(internalMutation, "#", "", 1)

	*self = UserID(internalMutation)
	return self
}

// ToUserMention Adds the necessary formatting for Discord to mention a User properly
func (self *UserID) ToUserMention() *UserID {
	internalMutation := string(*self)

	internalMutation = "<@" + internalMutation + ">"

	*self = UserID(internalMutation)
	return self
}

// ToRoleMention Adds the necessary formatting for Discord to mention a Role properly
func (self *UserID) ToRoleMention() *UserID {
	internalMutation := string(*self)

	internalMutation = "<@&" + internalMutation + ">"

	*self = UserID(internalMutation)
	return self
}

// ToChannelMention Adds the necessary formatting for Discord to mention a Channel properly
func (self *UserID) ToChannelMention() *UserID {
	internalMutation := string(*self)

	internalMutation = "<#" + internalMutation + ">"

	*self = UserID(internalMutation)
	return self
}

// ToString Converts to string for an easier usage.
func (self *UserID) ToString() string {
	return string(*self)
}
