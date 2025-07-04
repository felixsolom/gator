package commands

import (
	"context"
	"fmt"

	"github.com/felixsolom/gator/internal/database"
)

func HandlerFollowing(s *State, _ Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("couldn't get feeds for this user. Error: %w", err)
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("feed name: %v\n", feedFollow.FeedName)
	}
	return nil
}
