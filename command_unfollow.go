package main

import (
	"context"
	"fmt"

	"github.com/spollaL/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed for %s: %v", url, err)
	}

	deleteFollowParams := database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollowsForUser(context.Background(), deleteFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't delete follow for %s: %v", url, err)
	}

	fmt.Printf("Successfully delete follow for %s\n", url)
	return nil
}
