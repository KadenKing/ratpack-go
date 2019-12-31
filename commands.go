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

type pointsCommandGenerator func(command pointCommand, storage storage) func(sr slackRequest) error

type pointData struct {
	arguments string
	user      string
}

func newPointsCommandGenerator() pointsCommandGenerator {
	return func(command pointCommand, storage storage) func(sr slackRequest) error {
		switch command {
		case GIVE:
			return newGiveCommand(storage)
		default:
			return func(sr slackRequest) error {
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

func newGiveCommand(pi pointIncrementer) func(sr slackRequest) error {
	return func(sr slackRequest) error {
		pc := pointChange{UserChanging: sr.UserID}

		args := strings.Split(sr.Text, " ")
		if len(args) < 3 {
			return errors.New("Error: Too few arguments")
		}

		pc.User = args[0]

		points, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return fmt.Errorf("Could not parse point value as integer")
		}
		pc.Points = points

		message := getPointAnnotation(2, sr.Text)

		if message == "" {
			return errors.New("Error: reason for points is required")
		}

		pc.Reason = message

		err = pi.IncrementPoints(pc)
		return err
	}
}
