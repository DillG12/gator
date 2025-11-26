package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	// Delete all users from the database
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}

	fmt.Println("All users have been reset.")

	// Clear the current user in the config
	s.cfg.SetUser("")
	return nil
}
