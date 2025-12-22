package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get the feeds list: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.FeedName)
		fmt.Printf("* url: %s\n", feed.Url)
		fmt.Printf("* user: %s\n", feed.UserName)
	}
	return nil
}
