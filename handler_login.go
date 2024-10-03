package main

import (
	"context"
	"fmt"
)

func handlerLogin(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]

	_, err := state.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	err = state.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}