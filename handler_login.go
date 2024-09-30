package main

import "fmt"

func handlerLogin(state *state, cmd command) error {
	if state == nil || state.cfg == nil {
		return fmt.Errorf("no config state")
	}

	if len(cmd.args) == 0 {
		return fmt.Errorf("a username is required")
	}

	userName := cmd.args[0]
	err := state.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set: %s\n", userName)
	return nil
}