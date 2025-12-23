package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spollaL/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s Optional<limit>", cmd.Name)
	}

	limit := 2
	if len(cmd.Args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user %s: %v", user.Name, err)
	}

	for _, post := range posts {
		fmt.Println("----------------------")
		fmt.Printf("Title: %s\n", post.Title.String)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published: %s\n", post.PublishedAt)
		fmt.Println("----------------------")
	}

	return nil


}
