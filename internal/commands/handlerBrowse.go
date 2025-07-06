package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.Args) != 0 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil || limit <= 0 {
			return fmt.Errorf("argument must be a positive number: %w", err)
		}
	}

	limit32 := int32(limit)
	userID := database.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}

	params := database.GetPostsForUserParams{
		UserID: uuid.NullUUID(userID),
		Limit:  limit32,
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("cannot retrieve posts for this user: %w", err)
	}
	fmt.Println("=================================================")
	fmt.Println("================ printing posts ! ===============")
	for _, post := range posts {
		fmt.Println("------------------------------------")
		fmt.Printf("from: %v\n", post.Url)
		fmt.Printf("published: %v\n", post.PublishedAt)
		fmt.Printf("title: %v\n", post.Title.String)
		fmt.Printf("content: %v\n", post.Description.String)
	}
	return nil
}
