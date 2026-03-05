package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/buck/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("Insufficient Arguments")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	ctx := context.Background()

	db, err := s.db.Addfeed(ctx, database.AddfeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	create, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    db.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println(create)
	return nil
}

func Handlerfeeds(s *state, cmd command) error {
	db, err := s.db.Getfeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range db {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		user_name, err := s.db.GetId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(user_name)
	}
	return nil
}
