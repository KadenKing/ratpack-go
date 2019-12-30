package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *server) parseCommand(commandString string) (func() error, error) {
	args := strings.Split(commandString, " ")

	return newGiveCommand(s.storage, args)
}

func newGiveCommand(pi pointIncrementer, args []string) (func() error, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Give command expected 2 arguments, got %d", len(args))
	}

	points, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Could not parse point value as integer")
	}

	recipient := args[0]

	return func() error {
		err := pi.IncrementPoints(recipient, points)
		return err
	}, nil
}
