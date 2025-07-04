package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide a name and a link for the feed")
	}
	if len(cmd.Args) == 1 {
		return fmt.Errorf("please provide a link for the feed")
	}

	userID := database.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    uuid.NullUUID(userID),
	}
	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create new feed: %w", err)
	}
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Created at: %v\n", feed.CreatedAt)
	fmt.Printf("Updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("Url: %v\n", feed.Url)
	fmt.Printf("User's ID: %v\n", feed.UserID.UUID)

	feedID := database.NullUUID{
		UUID:  feed.ID,
		Valid: true,
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID(userID),
		FeedID:    uuid.NullUUID(feedID),
	}
	row, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't articulate user's followed feed. error: %w", err)
	}
	fmt.Printf("feed's name: %v\n", row.FeedName)
	fmt.Printf("user's name: %v\n", row.UserName)
	return nil
}
