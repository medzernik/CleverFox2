// Package command libraries for authenticating whether a user is an admin, owner, priviledged
package command

import (
	"CleverFox2/Info"
	"CleverFox2/logging"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// roleCheck checks the guild roles for the correct permissions and returns a map of all available roles
func roleCheckAdmin(s *discordgo.Session, guildID string) (map[string]struct{}, error) {
	roleOutput := make(map[string]struct{})

	guild, err := s.State.Guild(guildID)
	if err != nil {
		fmt.Println(err)
		return roleOutput, nil
	}

	for _, j := range guild.Roles {
		// This doesn't work properly. Only works when the server role has ADMINISTRATOR and NO OTHER priviledge set. Probably a bug in the library
		if j.Permissions == discordgo.PermissionAdministrator || j.Permissions == discordgo.PermissionAll {
			roleOutput[j.ID] = struct{}{}
		}
	}
	if len(roleOutput) == 0 {
		return roleOutput, errors.New("there are no roles present")
	} else {
		return roleOutput, nil
	}

}

// IsAdmin Backend function to check whether a user is admin
func IsAdmin(s *discordgo.Session, m *discordgo.InteractionCreate, user Info.UserID) (bool, error) {

	//get the guild pointer
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println(err)
		logging.Log.Warn("cannot find a single role within the guild\n", err)
		return false, err
	}

	//get all admin roles
	adminRoles, err := roleCheckAdmin(s, m.GuildID)
	if err != nil {
		fmt.Println(err)
		logging.Log.Warn(err)
		return false, err
	}

	//check whether a user has the admin privileges
	for _, Users := range guild.Members {
		//compare the userID we are looking for
		if Users.User.ID == user.ToString() {
			//get the user roles and check for each
			for _, rolesUser := range Users.Roles {
				//if the role exists in the admin roles map, if yes then return true
				if _, ok := adminRoles[rolesUser]; ok == true {
					return true, nil
				}
			}
		}
	}

	//else return false
	return false, nil

}

// IsOwner Checks if the userID called is an admin. Each command must call this check themselves
func IsOwner(s *discordgo.Session, m *discordgo.InteractionCreate, user Info.UserID) (bool, error) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return false, err
	}
	if user.ToString() == guild.OwnerID {
		return true, nil
	}
	return false, nil
}

func IsX(s *discordgo.Session, m *discordgo.InteractionCreate, cmd []interface{}) {

}

func isX(s *discordgo.Session, m *discordgo.InteractionCreate, cmd []interface{}) {

}
