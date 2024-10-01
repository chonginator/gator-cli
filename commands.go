package main

import (
	"errors"
)

type commands struct{
	commands map[string] func(*state, command) error
}

type command struct{
	name string
	args []string
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
