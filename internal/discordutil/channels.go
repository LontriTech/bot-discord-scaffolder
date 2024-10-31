package discordutil

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func FetchExistingChannels(dg *discordgo.Session, guildID string) (map[string]*discordgo.Channel, map[string]*discordgo.Channel, error) {
	guildChannels, err := dg.GuildChannels(guildID)
	if err != nil {
		return nil, nil, err
	}

	existingCategories := make(map[string]*discordgo.Channel)
	existingChannels := make(map[string]*discordgo.Channel)

	for _, channel := range guildChannels {
		if channel.Type == discordgo.ChannelTypeGuildCategory {
			nameKey := NormalizeName(channel.Name)
			existingCategories[nameKey] = channel
		} else {
			parentID := channel.ParentID
			nameKey := NormalizeName(channel.Name)
			key := fmt.Sprintf("%s|%s", parentID, nameKey)
			existingChannels[key] = channel
		}
	}
	return existingCategories, existingChannels, nil
}

func CreateCategory(session *discordgo.Session, guildID string, config CategoryConfig) (string, error) {
	var overwrites []*discordgo.PermissionOverwrite

	if config.Private {
		// Deny VIEW_CHANNEL for @everyone role
		overwrite := &discordgo.PermissionOverwrite{
			ID:   guildID, // The @everyone role ID is the same as the guild ID
			Type: discordgo.PermissionOverwriteTypeRole,
			Deny: discordgo.PermissionViewChannel,
		}
		overwrites = append(overwrites, overwrite)
	}

	var categoryName string

	if config.Prefix != "" {
		categoryName = AddNamePrefix(config.Prefix, config.Name)
	} else {
		categoryName = config.Name
	}

	channelData := discordgo.GuildChannelCreateData{
		Name:                 NormalizeNameConfig(categoryName),
		Type:                 discordgo.ChannelTypeGuildCategory,
		PermissionOverwrites: overwrites,
	}

	channel, err := session.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}

func CreateTextChannel(session *discordgo.Session, guildID, parentID string, config ChannelConfig) (string, error) {
	channelData := discordgo.GuildChannelCreateData{
		Name:     NormalizeName(config.Name),
		Type:     discordgo.ChannelTypeGuildText,
		ParentID: parentID,
		Topic:    config.Topic,
		NSFW:     config.NSFW,
		Position: config.Position,
	}

	channel, err := session.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}

func CreateVoiceChannel(session *discordgo.Session, guildID, parentID string, config ChannelConfig) (string, error) {
	channelData := discordgo.GuildChannelCreateData{
		Name:     NormalizeNameConfig(config.Name),
		Type:     discordgo.ChannelTypeGuildVoice,
		ParentID: parentID,
		Position: config.Position,
	}

	channel, err := session.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}

func CreateForumChannel(session *discordgo.Session, guildID, parentID string, config ChannelConfig) (string, error) {
	channelData := discordgo.GuildChannelCreateData{
		Name:     NormalizeNameConfig(config.Name),
		Type:     discordgo.ChannelTypeGuildForum,
		ParentID: parentID,
		Topic:    config.Topic,
		Position: config.Position,
	}

	channel, err := session.GuildChannelCreateComplex(guildID, channelData)
	if err != nil {
		return "", err
	}
	return channel.ID, nil
}
