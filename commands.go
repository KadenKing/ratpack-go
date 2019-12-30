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

type pointsCommandGenerator func(command pointCommand, storage storage) func(data pointData) error

type pointData struct {
	arguments string
	user      string
}

func newPointsCommandGenerator() pointsCommandGenerator {
	return func(command pointCommand, storage storage) func(data pointData) error {
		switch command {
		case GIVE:
			return newGiveCommand(storage)
		default:
			return func(data pointData) error {
				return errors.New("unsupported command")
			}

		}
	}
}

func getPointAnnotation(spaceNumber int, str string) string {
	for i := 0; i < spaceNumber; i++ {
		index := strings.Index(str, " ")
		str = str[index+1:]
	}
	return str
}

func newGiveCommand(pi pointIncrementer) func(data pointData) error {
	return func(data pointData) error {
		pc := pointChange{userChanging: data.user}

		args := strings.Split(data.arguments, " ")
		if len(args) < 3 {
			return errors.New("Error: Too few arguments")
		}

		pc.user = args[0]

		points, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return fmt.Errorf("Could not parse point value as integer")
		}
		pc.points = points

		message := getPointAnnotation(2, data.arguments)

		if message == "" {
			return errors.New("Error: reason for points is required")
		}

		pc.reason = message

		err = pi.IncrementPoints(pc)
		return err
	}
}
