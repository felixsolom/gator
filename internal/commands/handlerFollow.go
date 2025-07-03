package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide a valid link to a feed")
	}

	currentUser, err := s.Db.GetUser(context.Background(), s.PointerToConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("coudn't fetch current user %w", err)
	}

	feed, err := s.Db.FeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't geet feed by URL. error: %w", err)
	}

	feedID := database.NullUUID{
		UUID:  feed.ID,
		Valid: true,
	}

	userID := database.NullUUID{
		UUID:  currentUser.ID,
		Valid: true,
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID(userID),
		FeedID:    uuid.NullUUID(feedID),
	}
	row, err := s.Db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't articulate user's followed feed. error: %w", err)
	}
	fmt.Printf("feed's name: %v\n", row.FeedName)
	fmt.Printf("user's name: %v\n", row.UserName)
	return nil
}
