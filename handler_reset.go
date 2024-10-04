package main

import (
	"context"
	"fmt"
)

func handlerReset(state *state, cmd command) error {
	err := state.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Users reset successfully!")
	return nil
}