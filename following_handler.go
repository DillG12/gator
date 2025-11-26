package main

import (
	"context"
	"fmt"

	"github.com/DillG12/gator/internal/database"
)

func followingHandler(s *state, cmd command, user database.User) error {

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to retrieve followed feeds for user %s: %w", user.Name, err)
	}

	if len(followedFeeds) == 0 {
		fmt.Println("You are not following any feeds.")
	} else {
		fmt.Printf("Feeds %s is following:\n", user.Name)
		for _, feed := range followedFeeds {
			fmt.Printf("- %s\n", feed.FeedName)
		}
	}
	return nil
}
