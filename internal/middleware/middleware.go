package middleware

import (
	"context"
	"fmt"

	"github.com/felixsolom/gator/internal/commands"
	"github.com/felixsolom/gator/internal/database"
)

func MiddlewareLoggedIn(
	handler func(s *commands.State, cmd commands.Command, user database.User) error,
) func(*commands.State, commands.Command) error {
	return func(s *commands.State, cmd commands.Command) error {
		userName := s.PointerToConfig.CurrentUserName
		if userName == "" {
			return fmt.Errorf("user not logged in")
		}
		user, err := s.Db.GetUser(context.Background(), userName)
		if err != nil {
			return fmt.Errorf("user not found in database. error: %w", err)
		}
		return handler(s, cmd, user)
	}
}
