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
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: userName,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = state.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:		%v\n", user.Name)
}