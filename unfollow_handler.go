package main

import (
	"context"
	"fmt"

	"github.com/DillG12/gator/internal/database"
)

func unfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: unfollow <feed_url>")
	}
	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to retrieve feed by URL %s: %w", feedURL, err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed %s for user %s: %w", feedURL, user.Name, err)
	}

	fmt.Printf("User %s has unfollowed feed: %s\n", user.Name, feed.Name)
	return nil
}
