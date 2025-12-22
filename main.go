package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spollaL/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config %+v\n", cfg)
	s := state{config: &cfg}
	handlers := map[string]func(*state, command) error{}
	commands := commands{
		handlers: handlers,
	}
	commands.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("You must specify at least one argument")
	}
	commandName := args[1]
	if len(args) < 3 {
		log.Fatalf("You have to specify at least one argument after the %s command", commandName)
	}
	commandArg := args[2:]
	cmd := command{
		name: commandName,
		args: commandArg,
	}
	commands.run(&s, cmd)
}

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("Unkwown command")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
