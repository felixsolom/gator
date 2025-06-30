package commands

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, _ Command) error {
	feeds, err := s.Db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds, error: %w", err)
	}
	for _, feed := range feeds {
		fmt.Printf("user name: %v\n", feed.UserName)
		fmt.Printf("feed name: %v\n", feed.FeedName)
		fmt.Printf("feed url: %v\n", feed.FeedUrl)
	}
	return nil
}
