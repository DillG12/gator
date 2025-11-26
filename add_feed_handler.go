package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DillG12/gator/internal/database"
	"github.com/google/uuid"
)

func addFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	ctx := context.Background()

	feedID := uuid.New().ID()
	now := time.Now()

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        int32(feedID),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	feedFollowID := uuid.New().ID()
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        int32(feedFollowID),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed '%s' added successfully for user '%s'.\n", name, user.Name)
	fmt.Println(feed)

	return nil
}
