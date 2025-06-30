package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/felixsolom/gator/internal/commands"
	"github.com/felixsolom/gator/internal/config"
	"github.com/felixsolom/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	cfg, err := config.Read(".gatorconfig.json", &config.Config{})
	if err != nil {
		fmt.Println("Warning: using default config", err)
		cfg = &config.Config{}
	}

	if cfg.DbURL == "" {
		cfg.DbURL = database.DbURL + "?sslmode=disable"
	}
	fmt.Println("Using connection string:", cfg.DbURL)
	//connStr := "postgres://felixsolomon:@localhost:5432/gator?sslmode=disable"
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("Database connection failed", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	state := &commands.State{
		Db:              dbQueries,
		PointerToConfig: cfg,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Command name is needed to proceed")
		os.Exit(1)
	}

	commandsStruct := commands.NewCommandsStruct()
	commandsStruct.Register("register", commands.HandlerRegister)
	commandsStruct.Register("login", commands.HandlerLogin)
	commandsStruct.Register("reset", commands.HandlerResetAll)
	commandsStruct.Register("users", commands.HandlerUsers)
	commandsStruct.Register("agg", commands.HandlerAgg)
	commandsStruct.Register("addfeed", commands.HandlerAddFeed)
	commandsStruct.Register("feeds", commands.HandlerFeeds)
	fmt.Printf("Registered commands: %v\n", commandsStruct.Mapped)

	commandName := args[1]
	commandArgs := args[2:]
	currentCommand := commands.Command{
		Name: commandName,
		Args: commandArgs,
	}
	if err := commandsStruct.Run(state, currentCommand); err != nil {
		fmt.Printf("error executing %s: %v\n", commandName, err)
		os.Exit(1)
	}
	if err := config.Write(".gatorconfig.json", state.PointerToConfig); err != nil {
		fmt.Printf("error writing to json file, error: %v", err)
	}
}
