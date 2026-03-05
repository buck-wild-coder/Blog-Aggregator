package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/buck/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("not sufficient")
	}
	url := cmd.Args[0]
	ctx := context.Background()
	fmt.Println(url)
	feed, err := s.db.Getfeedurl(ctx, url)
	if err != nil {
		return err
	}
	cff := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	db, err := s.db.CreateFeedFollow(ctx, cff)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %s\n", db.FeedName)
	fmt.Printf("User: %s\n", db.UserName)
	return nil
}

func Handlerfollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	db, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}
	fmt.Println(db)
	return nil
}
