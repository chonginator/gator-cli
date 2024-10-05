package main

import (
	"context"
	"fmt"
	"html"
)

func handlerAggregate(state *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	for i, item := range feed.Channel.Items {
		feed.Channel.Items[i].Description = html.UnescapeString(item.Description)
		feed.Channel.Items[i].Title = html.UnescapeString(item.Title)
	}

	fmt.Printf("%+v", feed)
	return nil
}