package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(state *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name, url := cmd.args[0], cmd.args[1]

	user, err := state.db.GetUser(context.Background(), state.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user to add feed: %w", err)
	}

	feed, err := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add feed: %w", err)
	}

	fmt.Println("Feed created successfully!")
	printFeed(feed)
	 
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID: 			%v\n", feed.ID)
	fmt.Printf(" * Created: %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated: %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name: 		%v\n", feed.Name)
	fmt.Printf(" * URL: 		%v\n", feed.Url)
	fmt.Printf(" * UserID: 	%v\n", feed.UserID)
}