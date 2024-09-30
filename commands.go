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
	if c.commands == nil {
		c.commands = make(map[string]func(*state, command) error)
	}
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if len(c.commands) == 0 {
		return errors.New("no commands registered")
	}

	if s == nil || s.cfg == nil {
		return errors.New("no config state")
	}

	cmdFunc := c.commands[cmd.name]
	if cmdFunc == nil {
		return errors.New("couldn't retrieve handler for command")
	}

	err := cmdFunc(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
