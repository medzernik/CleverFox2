package passive

import (
	"CleverFox2/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

// DeleteCringe checks all channels for a specific gif from a specific person and then deletes it if it hasn't been posted in the GIF channel :)
//TODO: vycistit array po zmazani
//TODO: Checknut vsetky channely az na (blacklist)
func DeleteCringe(s *discordgo.Session) {
	var trackedIDTemp [1]string
	trackedIDTemp[0] = "206720832695828480"
	var trackedChannelTemp [1]string
	trackedChannelTemp[0] = "877632897006329896"

	for {
		time.Sleep(time.Second * 5)
		for k := range trackedChannelTemp {
			channel, _ := s.State.Channel(trackedChannelTemp[k])
			for i := range channel.Messages {
				fmt.Println(channel.Messages[i].Content)
				for j := range trackedIDTemp {
					if channel.Messages[i].Author.ID == trackedIDTemp[j] {
						if channel.Messages[i].Embeds != nil {
							for l := range channel.Messages[i].Embeds {
								if channel.Messages[i].Embeds[l].Type == discordgo.EmbedTypeGifv || strings.Contains(channel.Messages[i].Content, "gif") {
									fmt.Println(channel.Messages[i].ID)
									if _, err := s.ChannelMessageSend(config.Cfg.ChannelLog.CringeGIFChannelID, channel.Messages[i].Content); err != nil {
										s.ChannelMessageSend("877632897006329896", fmt.Sprintf("error removing message: "+err.Error()))
										fmt.Println("error removing message: " + err.Error())
									}
									if _, err := s.ChannelMessageSend("877632897006329896", "Kokotina je v <#893975128138973265>"); err != nil {
										s.ChannelMessageSend("877632897006329896", fmt.Sprintf("error removing message: "+err.Error()))
										fmt.Println("error removing message: " + err.Error())
									}
									if err := s.ChannelMessageDelete("877632897006329896", channel.Messages[i].ID); err != nil {
										s.ChannelMessageSend("877632897006329896", fmt.Sprintf("error removing message: "+err.Error()))
										fmt.Println("error removing message: " + err.Error())
									}

								}
							}
						}
					}
				}
			}
		}

	}
}
