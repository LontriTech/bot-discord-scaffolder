package util

import (
	"fmt"
	"os"
)

func GetEnv(key string, defaultValue ...string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	} else if len(defaultValue) > 0 {
		return defaultValue[0], nil
	}

	return "", fmt.Errorf("environment variable '%s' is not set and no default value was provided... ", key)
}
