package commands

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide valid credentials")
	}
	user := cmd.Args[0]
	_, err := s.Db.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("this user is not registered in database. error: %w", err)
	}
	s.PointerToConfig.SetUser(cmd.Args[0])
	fmt.Println("User was set")
	return nil
}
