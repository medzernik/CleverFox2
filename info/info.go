// Package info module to gather information about users, about servers, etc...
package info

import "github.com/bwmarrin/discordgo"

// UserServerInfo This function is able to list information about the user.
//Returns a struct to be able to work with it
//TODO: make it work with an appcommand
func UserServerInfo(s *discordgo.Session, i *discordgo.InteractionCreate, cmd []interface{}) {

}

// UserAdminInfo This function should be able to detect the information about a user by his ID
func UserAdminInfo(s *discordgo.Session, i *discordgo.InteractionCreate, cmd []interface{}) {

}

// GetServerInvite get the
func GetServerInvite(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
