package main

import (
	"context"
	"fmt"

	"github.com/chonginator/gator-cli/internal/database"
)

func middlewareLoggedIn(handler func(state *state, cmd command, user database.User) error) func(*state, command) error {
	return func(state *state, cmd command) error {
		user, err := state.db.GetUser(context.Background(), state.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get current user: %w", err)
		}

		err = handler(state, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}