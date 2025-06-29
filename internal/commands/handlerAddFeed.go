package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide a name and a link for the feed")
	}
	if len(cmd.Args) == 1 {
		return fmt.Errorf("please provide a link for the feed")
	}
	currentUser, err := s.Db.GetUser(context.Background(), s.PointerToConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("coudn't fetch current user %w", err)
	}
	userID := database.NullUUID{
		UUID:  currentUser.ID,
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
	return nil
}
