package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spollaL/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config %+v\n", cfg)
	programState := &state{config: &cfg}
	handlers := map[string]func(*state, command) error{}
	commands := commands{
		handlers: handlers,
	}
	commands.register("login", handlerLogin)
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
	commands.run(programState, cmd)
}
