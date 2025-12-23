package feedapi

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedUrl string) (RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("couldn't fetch the feed: %v", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("couldn't request to fetch the feed: %v", err)
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("couldn't read the fetched feed: %v", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(resData, &feed)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("couldn't decode the feed data: %v", err)
	}
	formatFeed(&feed)

	fmt.Printf("Successfully fetched feed for %s\n", feedUrl)
	return feed, nil
}

func formatFeed(feed *RSSFeed) {
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		item.Description = html.UnescapeString(item.Description)
		item.Title = html.UnescapeString(item.Title)
    feed.Channel.Item[i] = item
	}
}
