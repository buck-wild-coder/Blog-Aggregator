package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.name) < 2 {
		return errors.New("The login handler expects a single argument, the username!")
	}
	err := s.config.SetUser(cmd.argument[0])
	fmt.Println(s.config.CurrentUserName)
	if err != nil {
		return err
	}
	fmt.Println("The user has been set.")
	return nil
}
