package main

import (
	"context"
	"fmt"

	"github.com/spollaL/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("couldn't get the followed feed for user %s: %v", user.Name, err)
	}

	fmt.Printf("Feeds followed by %s\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}
