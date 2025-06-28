package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, _ Command) error {
	users, err := s.Db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("can't form users list. error: %w", err)
	}
	if len(users) == 0 {
		fmt.Println("no users currently in database")
		return nil
	}
	currentUser := s.PointerToConfig.CurrentUserName
	for _, user := range users {
		if user == currentUser {
			user = user + " (current)"
		}
		fmt.Printf("* %v\n", user)
	}
	return nil
}
