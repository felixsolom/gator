package commands

import (
	"context"
	"fmt"

	"github.com/felixsolom/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	rssFeed = rss.UnescapedRSS(rssFeed)
	fmt.Println(rssFeed)
	return nil
}
