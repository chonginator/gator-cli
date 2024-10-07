package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
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
			UserID: 	 user.ID,
			FeedID: 	 feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFollowing(state *state, cmd command, user database.User) error {
	feedFollows, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds current user is following: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user: %s\n", user.Name)
	for _, follow := range feedFollows {
		fmt.Printf(" * %s\n", follow.FeedName)
	}

	return nil
}

func handlerUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	feed, err := state.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = state.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User: %s\n", userName)
	fmt.Printf("* Feed: %s\n", feedName)
}