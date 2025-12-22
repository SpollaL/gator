package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spollaL/gator/internal/database"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Could get current users info: %v", err)
	}

	feedUrl := cmd.Args[1] 
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	}

	_, err = s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't created feed: %v", err)
	}
	
	fmt.Printf("Successfully added feed %s\n", feedUrl)
	return nil
}
