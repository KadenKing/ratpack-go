package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type pointCommand int

const (
	GIVE = iota
)

type pointsCommandGenerator func(command pointCommand, storage storage) func(arguments string) error

func newPointsCommandGenerator() pointsCommandGenerator {
	return func(command pointCommand, storage storage) func(arguments string) error {
		switch command {
		case GIVE:
			return newGiveCommand(storage)
		default:
			return func(arguments string) error {
				return errors.New("unsupported command")
			}

		}
	}
}

func newGiveCommand(pi pointIncrementer) func(arguments string) error {
	return func(arguments string) error {
		args := strings.Split(arguments, " ")
		if len(args) != 2 {
			return fmt.Errorf("Give command expected 2 arguments, got %d", len(args))
		}

		points, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return fmt.Errorf("Could not parse point value as integer")
		}

		recipient := args[0]
		err = pi.IncrementPoints(recipient, points)
		return err
	}
}
