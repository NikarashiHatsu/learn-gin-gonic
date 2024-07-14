package db_utils

import "os"

func SetEnvIfNotEmpty(envVar string, target *string) {
	value := os.Getenv(envVar)

	if value != "" {
		*target = value
	}
}
