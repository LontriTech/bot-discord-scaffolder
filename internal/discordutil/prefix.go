package discordutil

func AddNamePrefix(prefix string, name string) string {
	prefixedName := "[" + prefix + "] " + name
	return prefixedName
}
