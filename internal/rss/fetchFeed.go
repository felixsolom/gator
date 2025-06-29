package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("couldn't make a request: %w", err)
	}

	req.Header.Set("User-Agent", "Gator/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("couldn't get a response %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("couldn't read response body %w", err)
	}

	var rssFeed *RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("coudn't unmarshal body %w", err)
	}
	return rssFeed, nil
}

func UnescapedRSS(rss *RSSFeed) *RSSFeed {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for _, items := range rss.Channel.Item {
		items.Title = html.UnescapeString(items.Title)
		items.Description = html.UnescapeString(items.Description)
	}
	return rss
}
