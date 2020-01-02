package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type pointData struct {
	arguments string
	user      string
}

type whoDidWhat struct {
	Who     string `bson:"who"`
	Did     string `bson:"did"`
	Points  int64  `bson:"points"`
	ToWhom  string `bson:"towhom"`
	Because string `bson:"because"`
}

type whoDidWhatParser interface {
	Parse(sr slackRequest, idTranslater slackIDTranslater) (whoDidWhat, error)
}

type giveCommandParser struct{}

func newGiveCommandParserGenerator() func() giveCommandParser {
	return func() giveCommandParser {
		return giveCommandParser{}
	}
}

func (p giveCommandParser) Parse(sr slackRequest, idTranslater slackIDTranslater) (whoDidWhat, error) {
	did := "gave"

	args := strings.Split(sr.Text, " ")
	if len(args) < 3 {
		return whoDidWhat{}, errors.New("Error: Too few arguments")
	}

	toWhom := args[0]

	points, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return whoDidWhat{}, fmt.Errorf("Could not parse point value as integer")
	}

	because := getPointAnnotation(2, sr.Text)

	if because == "" {
		return whoDidWhat{}, errors.New("Error: reason for points is required")
	}

	who, err := idTranslater.GetProfileByID(sr.UserID)

	if err != nil {
		return whoDidWhat{}, nil
	}

	return whoDidWhat{
		who,
		did,
		points,
		toWhom,
		because,
	}, nil
}

func getPointAnnotation(spaceNumber int, str string) string {
	for i := 0; i < spaceNumber; i++ {
		index := strings.Index(str, " ")
		str = str[index+1:]
	}
	return str
}
