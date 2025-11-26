package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/DillG12/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUserName := s.cfg.CurrentUserName
		if currentUserName == "" {
			return errors.New("no user is currently logged in")
		}

		user, err := s.db.GetUser(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf("failed to retrieve logged-in user %s: %w", currentUserName, err)
		}

		return handler(s, cmd, user)
	}
}
