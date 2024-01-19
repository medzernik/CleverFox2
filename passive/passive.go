package passive

import (
	"CleverFox2/config"
	"CleverFox2/tviewsystem"
	"fmt"
	"net/url"
)

// replacer is a map that stores replacement values for specific keys.
// The keys in the map represent the strings to be replaced, while the
// corresponding values represent the replacement strings.
//
// Example:
// replacer["x.com"] = "vxtwitter.com"
// replacer["twitter.com"] = "vxtwitter.com"
//
// In the above example, whenever "x.com" or "twitter.com" is found in
// a given string, it should be replaced with "vxtwitter.com".
var replacer = map[string]string{
	"x.com":       "vxtwitter.com",
	"twitter.com": "vxtwitter.com",
}

func ReplaceURL(link *url.URL) {
	// Check if the subsystem is enabled
	if config.Cfg.Modules.LinkFixer == false {
		return
	}
	// Get the main link part

	if value, exists := replacer[link.Host]; exists == true {
		tviewsystem.MainViewPush(fmt.Sprintf("Replacing: %s with: %s", link.Host, value))
		link.Host = value
	}
	// Delete the spying part of URL?

}

/*
// DeleteCringe checks all channels for a specific gif from a specific person and then deletes it if it hasn't been posted in the GIF channel :)
func DeleteCringe(s *discordgo.Session) {
	//user to track and channels to ignore
	trackedID := []string{"206720832695828480"}
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
			channel.Messages = nil
		}

	}
}

*/
