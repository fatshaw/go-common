package go_common

import "os"

func GetEnvValueWithDefault(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		value = defaultValue
	}
	return value
}