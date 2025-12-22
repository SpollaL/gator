package main

import (
	"context"
	"fmt"

	"github.com/spollaL/gator/internal/feed"
)

func handlerAgg(_ *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error calling FetchFeed: %v", err)
	}
	fmt.Println("Feed for 'https://www.wagslane.dev/index.xml'")
	fmt.Print(feed)
	return nil
}
