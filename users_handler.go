package main

import (
	"context"
	"fmt"
)

func usersHandler(s *state, cmd command) error {
	if users, err := s.db.GetUsers(context.Background()); err != nil {
		return fmt.Errorf("failed to retrieve users: %w", err)
	} else if len(users) == 0 {
		fmt.Println("No registered users found.")
	} else {
		fmt.Println("Registered users:")
		for _, user := range users {
			if user.Name == s.cfg.CurrentUserName {
				fmt.Printf("* %s (current)\n", user.Name)
			} else {
				fmt.Printf("* %s\n", user.Name)
			}
		}
	}
	return nil
}
