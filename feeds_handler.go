package main

import (
	"context"
	"fmt"
)

func feedsHandler(s *state, cmd command) error {
	// Implementation for printing feeds by user will go here
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found for the current user.")
		return nil
	}
	for _, feed := range feeds {
		userName, err := s.db.GetUserNameByFeedID(ctx, feed.ID)
		if err != nil {
			return fmt.Errorf("failed to retrieve user name for feed %d: %w", feed.ID, err)
		}
		fmt.Printf("Feed: %s, URL: %s, User: %s\n", feed.Name, feed.Url, userName)
	}
	return nil
}
