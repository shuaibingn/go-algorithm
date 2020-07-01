package utils

import "os"

func GetEnv(opts ...string) string {
	switch len(opts) {
	case 1:
		key := opts[0]
		return os.Getenv(key)
	case 2:
		key, defaultValue := opts[0], opts[1]
		value := os.Getenv(key)
		if value == "" {
			return defaultValue
		} else {
			return value
		}
	default:
		return ""
	}
}
