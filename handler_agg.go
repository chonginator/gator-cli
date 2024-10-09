package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerAggregate(state *state, cmd command, user database.User) error {
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
		scrapeFeeds(state)
	}
}

func scrapeFeeds(state *state) {
	nextFeed, err := state.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")

	err = state.db.MarkFeedFetched(context.Background(), nextFeed.ID)
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
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse post %s publication date: %v", item.Title, item.PubDate)
		}
		_, err = state.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: item.Description != "",
			},
			PublishedAt: pubDate,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			log.Printf("Couldn't add post %s to database: %v", item.Title, err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(rssFeed.Channel.Item))
}