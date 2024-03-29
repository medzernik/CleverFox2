// Package config Configuration module that holds the configuration logic.
// This is the global settings for the bot.
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// Config Configuration type structure to use in memory.
type Config struct {
	ServerInfo struct {
		WeatherAPIKey string `yaml:"weatherAPIKey"`
		ServerToken   string `yaml:"serverToken"`
		BotLogo       string `yaml:"botLogo"`
		BotName       string `yaml:"botName"`
		LogLevel      string `yaml:"logLevel"`
	} `yaml:"serverInfo"`
	Modules struct {
		LinkFixer bool `yaml:"linkFixer"`
	} `yaml:"modules"`
	ChannelLog struct {
		ChannelLogID       string `yaml:"channelLogID"`
		CringeGIFChannelID string `yaml:"cringeGIFChannelID"`
	} `yaml:"channelLog"`
	Tracked struct {
		TrackedID []string `yaml:"trackedID"`
	}
}

var Cfg Config

// LoadConfig Loads the config file. It must be in the root of the directory, next to the main executable.
func LoadConfig() {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Config loaded.")

}

// SaveConfig This function allows you to save the current config to the file
func SaveConfig() {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(&Cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Configure This function allows you to set values in the config straight from discord.
func Configure() {

}
