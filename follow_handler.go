package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DillG12/gator/internal/database"
	"github.com/google/uuid"
)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("feed URL argument is required")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to retrieve feed by URL %s: %w", feedURL, err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        int32(uuid.New().ID()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow for user %s and feed %s: %w", user.Name, feedURL, err)
	}

	fmt.Printf("User %s is now following feed: %s\n", user.Name, feed.Name)
	return nil
}
