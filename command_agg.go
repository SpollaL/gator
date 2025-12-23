package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spollaL/gator/internal/database"
	"github.com/spollaL/gator/internal/feedapi"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting Feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)	
		fmt.Println(err)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't collect next feed to fetch: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now()},
		ID: feed.ID},
	)
	if err != nil {
		return fmt.Errorf("couldn't mark feed fetched: %v", err)
	}

	feedData, err := feedapi.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %v", err)
	}
	
	for _, item := range feedData.Channel.Item {

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("couldn't parse date: %s\n", item.PubDate)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{String: item.Title, Valid: item.Title != ""},
			Url: item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubDate,
			FeedID: feed.ID,
		})

		if err != nil {
			fmt.Printf("couldn't create post for %s: %v\n", feed.Url, err)
		} else {
			fmt.Printf("Successfully store post for feed %s\n", feed.Url)
		}
	}
	return nil
}
