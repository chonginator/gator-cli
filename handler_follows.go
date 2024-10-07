package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	currentUser, err := state.db.GetUser(context.Background(), state.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, err := state.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollow, err := state.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID: 			 uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID: 	 currentUser.ID,
			FeedID: 	 feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFollowing(state *state, cmd command) error {
	currentUser, err := state.db.GetUser(context.Background(),
		state.cfg.CurrentUserName,
	)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}
	
	feedFollows, err := state.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds current user is following: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user: %s\n", currentUser.Name)
	for _, follow := range feedFollows {
		fmt.Printf(" * %s\n", follow.FeedName)
	}

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User: %s\n", userName)
	fmt.Printf("* Feed: %s\n", feedName)
}