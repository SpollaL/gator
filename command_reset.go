package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset the db: %v", err)
	}

	fmt.Println("db reset was successful!")
	return nil
}
