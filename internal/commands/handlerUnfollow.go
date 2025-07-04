package commands

import (
	"context"
	"fmt"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide feed's URL to unfollow")
	}
	feed, err := s.Db.FeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't resolve url to feed info. error: %w", err)
	}

	feedID := database.NullUUID{
		UUID:  feed.ID,
		Valid: true,
	}

	userID := database.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	params := database.DeleteFeedFollowParams{
		UserID: uuid.NullUUID(userID),
		FeedID: uuid.NullUUID(feedID),
	}
	if err := s.Db.DeleteFeedFollow(context.Background(), params); err != nil {
		return fmt.Errorf("couldn't delete feed from record. error: %w", err)
	}
	return nil
}
