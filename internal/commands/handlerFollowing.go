package commands

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, _ Command) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), s.PointerToConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get feeds for this user. Error: %w", err)
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("feed name: %v", feedFollow.FeedName)
	}
	return nil
}
