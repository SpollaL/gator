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
	fmt.Printf("Read config %+v\n", cfg)

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

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}
	commandName := args[1]
	if len(args) < 3 {
		log.Fatalf("You have to specify at least one argument after the %s command", commandName)
	}
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




