package tviewsystem

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rivo/tview"
)

func UpdateMembers(s *discordgo.Session) {
	var memberList []string
	var nickList []string

	for _, guild := range s.State.Guilds {
		for _, member := range guild.Members {
			memberList = append(memberList, member.User.Username)
			if member.Nick == "" {
				nickList = append(nickList, "---")
			} else {
				nickList = append(nickList, member.Nick)
			}
		}

	}
	for _, test := range memberList {
		MainViewPush(fmt.Sprintf(test))
	}
	MemberListPush(memberList, nickList)
}

func UserAction() {
	dropdown := tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"Info", "Mute", "Kick", "Ban"}, nil)

	app.SetFocus(dropdown).Run()

}
