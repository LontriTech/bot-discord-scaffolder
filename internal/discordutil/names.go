package discordutil

import (
	"fmt"

	"bot-discord-scaffolder/internal/util"
)

func NormalizeChannelName(name string, channelType ...string) string {
	if len(channelType) == 1 && channelType[0] == "text" {
		name = util.FullySanatize(name)
	} else {
		name = util.RemoveExtraWhitespace(name)
	}

	return name
}

func NormalizeCategoryName(name string, prefix ...string) string {
	if len(prefix) == 1 {
		name = fmt.Sprintf("[%s] %s", prefix[0], name)
	}

	name = util.RemoveExtraWhitespace(name)

	return name
}

func NormalizeConfigKeyName(name string) string {
	name = util.FullySanatize(name)

	return name
}
