package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/felixsolom/gator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func ScrapeFeeds(s *State, cmd Command, userID uuid.NullUUID) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background(), userID)
	if err == sql.ErrNoRows {
		fmt.Println("No feeds left to fetch")
	}
	if err != nil {
		return fmt.Errorf("couldn't fetch the next feed: %w", err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("Failed to fetch or parse RSS: %v\n", err)
		return fmt.Errorf("couldn't fetch the rss feed: %w", err)
	}
	fmt.Printf("Fetched %d items from %s\n", len(rssFeed.Channel.Item), nextFeed.Name)

	for _, item := range rssFeed.Channel.Item {

		pubDate := time.Now()
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			pubDate = t
		} else if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			pubDate = t
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{String: item.Title, Valid: item.Title != ""},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}
		fmt.Printf("Title:  %v\n", postParams.Title)
		fmt.Printf("Url:  %v\n", postParams.Url)

		_, err = s.Db.CreatePost(context.Background(), postParams)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				continue
			}
			fmt.Printf("couldn't insert post: %v\n", err)
			continue
		}
	}
	fmt.Printf("About to mark feed as fetched: %s (%s)\n", nextFeed.Name, nextFeed.ID)

	params := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}

	_, err = s.Db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		fmt.Printf("Failed to mark feed as fetched: %v\n", err)
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}
	fmt.Printf("Successfully marked feed as fetched: %s (%s)\n", nextFeed.Name, nextFeed.ID)
	return nil
}
