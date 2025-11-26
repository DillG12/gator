package main

import (
	"context"
	"errors"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username argument is required")
	}

	username := cmd.args[0]

	if _, err := s.db.GetUser(context.Background(), username); err != nil {
		return fmt.Errorf("user %s does not exist: %w", username, err)
	}

	s.cfg.SetUser(username)

	fmt.Printf("Logged in as user: %s\n", username)
	return nil
}
