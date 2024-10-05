package main

import (
	"context"
	"fmt"
)

func handlerAggregate(state *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v", feed)
	return nil
}
