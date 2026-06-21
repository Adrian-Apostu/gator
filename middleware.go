package main

import (
	"context"
	"errors"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return errors.New("no user logged in")
		}

		return handler(s, c, user)
	}
}
