package main

import (
	"github.com/DillG12/gator/internal/config"
	"github.com/DillG12/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
