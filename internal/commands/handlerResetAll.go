package commands

import (
	"context"
	"fmt"
)

func HandlerResetAll(s *State, _ Command) error {
	if err := s.Db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("unable to reset database %w", err)
	}
	fmt.Println("database was reset")
	return nil
}
