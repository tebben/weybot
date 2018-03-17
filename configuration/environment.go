package configuration

import (
	"os"
)

// SetEnvironmentVariables changes config settings when certain environment variables are found
func SetEnvironmentVariables(conf *Config) {
	token := os.Getenv("WEYBOT_TOKEN")
	if token != "" {
		conf.Token = token
	}
}
