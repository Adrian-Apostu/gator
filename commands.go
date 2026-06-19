package main

import (
	"errors"
)

var handlers = map[string]func(*state, command) error{
	"login":    handlerLogin,
	"register": handlerRegister,
	"reset":    handlerReset,
	"users":    handlerUsers,
	"agg":      handlerAgg,
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
