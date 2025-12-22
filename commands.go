package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
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
