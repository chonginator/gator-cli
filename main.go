package main

import (
	"log"
	"os"

	"github.com/chonginator/gator-cli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments provided")
	}
	cmdName := os.Args[1]

	cmdArgs := []string{}
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = commands.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}