package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DillG12/gator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = specifiedLimit
		}
	}

	// Retrieve posts for the current user with the specified limit
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve posts for user %s: %w", s.cfg.CurrentUserName, err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found for the current user.")
		return nil
	}

	fmt.Printf("Posts for user %s:\n", s.cfg.CurrentUserName)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("2006-01-02 15:04:05"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
