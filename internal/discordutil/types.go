package discordutil

type ChannelConfig struct {
	Type     string `yaml:"type"`
	Name     string `yaml:"name"`
	Topic    string `yaml:"topic,omitempty"`
	NSFW     bool   `yaml:"nsfw,omitempty"`
	Position int    `yaml:"position,omitempty"`
}

type CategoryConfig struct {
	Name     string          `yaml:"name"`
	Prefix   string          `yaml:"prefix,omitempty"`
	Private  bool            `yaml:"private,omitempty"`
	Channels []ChannelConfig `yaml:"channels"`
}

type Config struct {
	Categories []CategoryConfig `yaml:"categories"`
}
