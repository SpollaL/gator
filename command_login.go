package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return errors.New("login command expects a single username")
	}

	userName := cmd.args[0]
	err := s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set: %s\n", userName)
	return nil
}
