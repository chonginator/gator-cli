package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
)

func handlerAggregate(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration string: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(state.db)
	}
}

func scrapeFeeds(db *database.Queries) {
	nextFeed, err := db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")

	err = db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("Couldn't mark next feed %s fetched: %v", nextFeed.Name, err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", nextFeed.Name, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(rssFeed.Channel.Item))
}