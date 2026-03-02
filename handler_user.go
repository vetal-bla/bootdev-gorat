package main

import (
	"context"
	// "errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vetal-bla/bootdev-gorat/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println("No user in database")
		return err
	}

	err = s.state.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s switched successfully", username)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	newUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	err = s.state.SetUser(username)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User created:\nID: %s\ncreated_at: %v\nname: %s\n", newUser.ID, newUser.CreatedAt, newUser.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("there no additional arguments for that command")
	}
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Can't remove users %w", err)
	}

	fmt.Println("users table was cleared")

	return nil
}

func handlerGetUsres(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("there no additional arguments for that command")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("cant get all users from database")
	}

	currentUser := s.state.CurrentUserName
	for _, u := range users {
		displayString := fmt.Sprintf("* %s", u.Name)
		if u.Name == currentUser {
			displayString = displayString + " (current)"
		}
		fmt.Println(displayString)
	}

	return nil
}
