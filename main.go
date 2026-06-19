package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)
import (
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		commands: make(map[string]func(*state, command) error),
	}

	for name, handler := range handlers {
		cmds.register(name, handler)
	}

	err = cmds.run(&s, command{name: args[0], args: args[1:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
