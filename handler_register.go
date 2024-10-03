package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chonginator/gator-cli/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]

	user, err := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: userName,
	})
	if err != nil {
		return err
	}

	err = state.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User created: %+v\n", user)

	return nil
}