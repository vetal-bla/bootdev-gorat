package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vetal-bla/bootdev-gorat/internal/database"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("there no additional arguments for that command")
	}
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Can't get rss feed %w", err)
	}

	fmt.Println(rss)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: %s <title> <url>", cmd.Name)
	}

	title := cmd.Args[0]
	urlFeed := cmd.Args[1]
	currentUser, err := s.db.GetUser(context.Background(), s.state.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Cant get erros from database: %s, %v", s.state.CurrentUserName, err)
	}

	feed := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       urlFeed,
		UserID:    currentUser.ID,
	}

	newFeed, err := s.db.AddFeed(context.Background(), feed)
	if err != nil {
		fmt.Errorf("Cant add feed. Sql error: %v", err)
	}
	if newFeed.Url == "" {
		fmt.Printf("Url %s alread exist in database", urlFeed)
	} else {
		fmt.Printf("ID: %s\n", newFeed.ID)
		fmt.Printf("Name (title): %s\n", newFeed.Name)
		fmt.Printf("URL: %s\n", newFeed.Url)
		fmt.Printf("User id: %s\n", newFeed.UserID)
	}

	return nil
}
