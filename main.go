package main

import (
	"fmt"
	"log"

	"github.com/DillG12/gator/gator_internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// Example usage of the Config struct
	cfg.SetUser("dill")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fmt.Printf("%v\n", cfg)
}
