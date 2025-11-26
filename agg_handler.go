package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	. "github.com/DillG12/gator/internal/database"
	"github.com/DillG12/gator/internal/rss"
)

func aggHandler(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: agg <time_between_requests>")
	}

	time_between_requests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time duration %s: %w", cmd.args[0], err)
	}

	fmt.Printf("Starting feed aggregation every %s...\n", time_between_requests)

	ticker := time.NewTicker(time_between_requests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		fmt.Println("Fetching feeds...\n")
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error during feed scraping: %v\n", err)
		} else {
			fmt.Println("Feed scraping completed successfully.\n")
		}
	}
	return nil
}

func testFetchHandler(s *state, cmd command) error {

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to parse the xml: %w", err)
	}

	fmt.Println(feed)
	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background(), sql.NullTime{Time: time.Now(), Valid: true})
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch: %w", err)
	}
	s.db.MarkFeedAsFetched(context.Background(), MarkFeedAsFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
	})

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	// Insert posts into the database

	for _, item := range feedData.Channel.Item {
		publishedAt := time.Now()
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = t
		}
		_, err := s.db.CreatePost(context.Background(), CreatePostParams{
			ID:          int32(uuid.New().ID()),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Fatalf("failed to create post for item %s: %v\n", item.Title, err)
		} else {
			fmt.Printf("Created post for item %s\n", item.Title)
		}
	}

	return nil
}
