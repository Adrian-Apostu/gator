package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

type state struct {
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

	s := state{cfg: &cfg}

	cmds := &commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	err = cmds.run(&s, command{name: args[0], args: args[1:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
