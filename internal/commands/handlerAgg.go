package commands

import (
	"fmt"
	"time"

	"github.com/felixsolom/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please provide time period between requests")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("please provide valid time format: %w", err)
	}
	fmt.Println("======+ RUNNING AGGREGATOR +======")
	fmt.Println("------ Press Ctrl-C to quit ------")
	ticker := time.NewTicker(timeBetweenRequests)
	userID := database.NullUUID{
		UUID:  user.ID,
		Valid: true,
	}
	for ; ; <-ticker.C {
		ScrapeFeeds(s, cmd, uuid.NullUUID(userID))
	}
}
