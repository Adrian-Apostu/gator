package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("Please provide a name and url.")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("Failed to get the current user!")
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return errors.New("Failed to add feed to the current user!")
	}

	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return errors.New("Could not fetch feeds!")
	}

	for _, feed := range feeds {
		author, err := s.db.GetUsersById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("* Name:        %s\n", feed.Name)
		fmt.Printf("* URL:         %s\n", feed.Url)
		fmt.Printf("* User:        %s\n", author.Name)
		fmt.Println("----------------------------------------")
	}
	return nil
}
