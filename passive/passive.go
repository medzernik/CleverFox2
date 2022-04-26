package passive

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

func DeleteCringe(s *discordgo.Session) {
	for {
		channel, _ := s.State.Channel("877632897006329896")
		for i := range channel.Messages {
			fmt.Println(channel.Messages[i].Content)
		}
		time.Sleep(time.Second * 1)
	}
}
