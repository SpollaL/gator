package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spollaL/gator/internal/config"
	"github.com/spollaL/gator/internal/database"
)

type state struct {
	db *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Successfully read config\n")

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	dbQueries := database.New(db)
	programState := &state{
		db: dbQueries,
		config: &cfg,
	}

	commands := commands{
		handlers: map[string]func(*state, command) error{},
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)
	commands.register("follow", handlerFollow)
	commands.register("following", handlerFollowing)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}
	commandName := args[1]
	commandArg := args[2:]
	cmd := command{
		Name: commandName,
		Args: commandArg,
	}
	err = commands.run(programState, cmd)
	if err != nil {
		log.Fatalf("command %s exited with error: %v", cmd.Name, err)
	}
}
