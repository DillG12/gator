package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DillG12/gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("must provide a username to register")
	}

	username := cmd.args[0]
	if existingUser, err := s.db.GetUser(context.Background(), username); err == nil {
		return fmt.Errorf("username %s is already in database: %v", username, existingUser)
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        int32(uuid.New().ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	s.cfg.SetUser(username)

	fmt.Printf("Registered user: %s\n", user.Name)
	return nil
}
