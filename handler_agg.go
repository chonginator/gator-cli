package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
)

func handlerAggregate(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse time_between_reqs duration string: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(state.db)
		if err != nil {
			return fmt.Errorf("couldn't scrape feed: %w", err)
		}
	}
}

func scrapeFeeds(db *database.Queries) error {
	nextFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}

	err = db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
	})
	if err != nil {
		return fmt.Errorf("couldn't mark next feed as fetched: %w", err)
	}
	fmt.Printf("%s feed successfully marked as fetched!\n", nextFeed.Name)

	fmt.Printf("fetching %s...", nextFeed.Url)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	}
	fmt.Printf("%s fetched successfully!\n", nextFeed.Url)

	fmt.Println("Posts:")
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

	return nil
}