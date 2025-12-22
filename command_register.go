package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spollaL/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	userName := cmd.Args[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	s.config.SetUser(userName)
	fmt.Printf("User %s was created with data: %+v\n", userName, user)
	return nil
}
