package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("command name is needed")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create new user: %v", err)
	}
	fmt.Println("New user created!")
	fmt.Printf("User id: %v\n", user.ID)
	fmt.Printf("User created at: %v\n", user.CreatedAt)
	fmt.Printf("User name: %v\n", user.Name)

	return nil
}
