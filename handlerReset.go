package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())

	if err != nil {
		return errors.New("Failed to reset users!")
	}
	fmt.Println("Users database reset successfully!")
	return nil
}
