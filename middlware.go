package main

import (
	"context"
	"fmt"

	"github.com/vetal-bla/bootdev-gorat/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.state.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Cant get user form db: %v", err)
		}
		return handler(s, cmd, user)
	}
}
