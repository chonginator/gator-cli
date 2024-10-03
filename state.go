package main

import (
	"github.com/chonginator/gator-cli/internal/config"
	"github.com/chonginator/gator-cli/internal/database"
)

type state struct{
	db *database.Queries
	cfg *config.Config
}
