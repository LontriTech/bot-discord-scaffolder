package main

import (
	"fmt"
	"log"

	"bot-discord-scaffolder/internal/discordutil"
	"bot-discord-scaffolder/internal/util"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, tokenErr := util.GetEnv("SCAFFOLDER_DISCORD_BOT_TOKEN")
	guildID, guildIDErr := util.GetEnv("SCAFFOLDER_DISCORD_GUILD_ID")
	configFile, configFileErr := util.GetEnv("SCAFFOLDER_CONFIG_FILE", ".discord-scaffolder.yaml")

	if tokenErr != nil {
		log.Fatalln(tokenErr)
	}

	if guildIDErr != nil {
		log.Fatalln(guildIDErr)
	}

	if configFileErr != nil {
		log.Fatalln(configFileErr)
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v \n", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuilds

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v \n", err)
	}
	defer dg.Close()

	fmt.Println("Bot is now running... Press CTRL+C to exit. ")

	config, err := discordutil.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading configuration: %v \n", err)
	}

	existingCategories, existingChannels, err := discordutil.FetchExistingChannels(dg, guildID)
	if err != nil {
		log.Fatalf("Error fetching existing channels: %v \n", err)
	}

	for _, categoryConfig := range config.Categories {
		var categoryID string
		var categoryKey string
		var categoryName string

		categoryName = discordutil.NormalizeCategoryName(categoryConfig.Name, categoryConfig.Prefix)

		categoryKey = discordutil.NormalizeConfigKeyName(categoryName)

		if existingCategory, exists := existingCategories[categoryKey]; exists {
			categoryID = existingCategory.ID
			fmt.Printf("Category: '%s' already exists with ID: '%s'. Skipping creation... \n", categoryName, categoryID)
		} else {
			categoryID, err = discordutil.CreateCategory(dg, guildID, categoryConfig, categoryName)
			if err != nil {
				log.Printf("Error creating category '%s': %v", categoryName, err)
				continue
			}
			fmt.Printf("Created category: '%s' with ID '%s' \n", categoryName, categoryID)

			existingCategories[categoryKey] = &discordgo.Channel{
				ID:   categoryID,
				Name: categoryName,
			}
		}

		for _, channelConfig := range categoryConfig.Channels {
			var channelID string
			channelKey := discordutil.NormalizeConfigKeyName(channelConfig.Name)
			channelName := discordutil.NormalizeChannelName(channelConfig.Name, channelConfig.Type)

			key := fmt.Sprintf("%s|%s", categoryID, channelKey)

			if existingChannel, exists := existingChannels[key]; exists {
				channelID = existingChannel.ID
				fmt.Printf("Channel: '%s' under Category: '%s' already exists with ID: '%s'. Skipping creation... \n", channelName, categoryName, channelID)
			} else {
				switch channelConfig.Type {
				case "text":
					channelID, err = discordutil.CreateTextChannel(dg, guildID, categoryID, channelConfig, channelName)
				case "voice":
					channelID, err = discordutil.CreateVoiceChannel(dg, guildID, categoryID, channelConfig, channelName)
				case "forum":
					channelID, err = discordutil.CreateForumChannel(dg, guildID, categoryID, channelConfig, channelName)
				default:
					log.Printf("Unknown channel type: '%s' for channel: '%s' \n", channelConfig.Type, channelName)
					continue
				}

				if err != nil {
					log.Printf("Error creating Channel '%s': %v \n", channelName, err)
					continue
				}
				fmt.Printf("Created Channel: '%s'\n - Type: '%s'\n - Category: '%s'\n - ID: '%s' \n", channelName, channelConfig.Type, categoryName, channelID)
				existingChannels[key] = &discordgo.Channel{
					ID:       channelID,
					Name:     channelConfig.Name,
					ParentID: categoryID,
				}
			}
		}
	}

	fmt.Println("Shutting down bot... ")
}
