package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/chonginator/gator-cli/internal/config"
	"github.com/chonginator/gator-cli/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		db: dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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