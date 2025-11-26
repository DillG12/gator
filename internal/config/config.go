package config

import (
	"log"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		// Handle the error appropriately, for example:
		log.Printf("Failed to write config: %v", err)
	}
}
