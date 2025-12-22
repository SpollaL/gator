package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	err = s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set: %s\n", userName)
	return nil
}
