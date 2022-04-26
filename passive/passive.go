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
	//user to track and channels to ignore
	trackedID := []string{"272292297020932096"}
	trackedChannelBlacklist := []string{"893975128138973265", "893231726267101185"}

	guild, err := s.State.Guild("869556493316395091")
	if err != nil {
		s.ChannelMessageSend(config.Cfg.ChannelLog.ChannelLogID, fmt.Sprintf("Error getting guild, disabling gif checking", err))
		return
	}

	channelsTemp := guild.Channels[:]
	var channelsCheckID []string

	//find all the channels that aren't blacklisted and get their IDs
	for i := range channelsTemp {
		for j := range trackedChannelBlacklist {
			if channelsTemp[i].ID != trackedChannelBlacklist[j] {
				channelsCheckID = append(channelsCheckID, channelsTemp[i].ID)
			}
		}
	}

	for {
		time.Sleep(time.Second * 5)

		for k := range channelsCheckID {
			channel, _ := s.State.Channel(channelsCheckID[k])
			for i := range channel.Messages {
				for j := range trackedID {
					if channel.Messages[i].Author.ID == trackedID[j] {
						if channel.Messages[i].Embeds != nil {
							for l := range channel.Messages[i].Embeds {
								if channel.Messages[i].Embeds[l].Type == discordgo.EmbedTypeGifv || strings.Contains(channel.Messages[i].Content, "gif") {
									fmt.Println("Found gif from X in: ", channel.Messages[i].ID)
									s.ChannelMessageSend(config.Cfg.ChannelLog.CringeGIFChannelID, channel.Messages[i].Content)
									s.ChannelMessageSend(channel.Messages[i].ChannelID, "Kokotina je v <#893975128138973265>")

									if err := s.ChannelMessageDelete(channel.Messages[i].ChannelID, channel.Messages[i].ID); err != nil {
										s.ChannelMessageSend(channel.Messages[i].ID, fmt.Sprintf("error removing message: "+err.Error()))
										fmt.Println("error removing message: " + err.Error())
									}
								}
							}
						}
					}

				}

			}
			channel.Messages = channel.Messages[:0]
		}

	}
}
