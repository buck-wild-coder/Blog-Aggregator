package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/buck/gator/internal/database"
	"github.com/google/uuid"
)

func addfeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("Insufficient Arguments")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	ctx := context.Background()
	getid, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	db, err := s.db.Addfeed(ctx, database.AddfeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    getid.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(db)
	return nil
}
