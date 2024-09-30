package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chonginator/gator-cli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)

	s := state{
		cfg: &cfg,
	}

	commands := commands{
		commands: map[string]func(*state, command) error {},
	}

	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("not enough arguments provided")
	}
	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = commands.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}