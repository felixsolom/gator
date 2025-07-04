package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/felixsolom/gator/internal/rss"
	"github.com/google/uuid"
)

func ScrapeFeeds(s *State, cmd Command, userID uuid.NullUUID) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background(), userID)
	if err != nil {
		return fmt.Errorf("couldn't fetch the next feed: %w", err)
	}
	params := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}
	_, err = s.Db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}
	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch the rss feed: %w", err)
	}
	for _, items := range rssFeed.Channel.Item {
		fmt.Printf("--> %v\n", items.Title)
	}
	return nil
}
