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
	Who      string `bson:"who"`
	WhoID    string `bson:"who_id"`
	Did      string `bson:"did"`
	Points   int64  `bson:"points"`
	ToWhom   string `bson:"towhom"`
	ToWhomID string `bson:"towhom_id"`
	Because  string `bson:"because"`
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

	because := parseReasonAfterSpace(2, sr.Text)

	who, err := idTranslater.GetProfileByID(sr.UserID)
	if err != nil {
		return whoDidWhat{}, err
	}

	whoID := sr.UserID

	toWhomProfile, err := idTranslater.GetProfileByUsername(toWhom)

	if err != nil {
		return whoDidWhat{}, err
	}

	return whoDidWhat{
		Who:      who,
		WhoID:    whoID,
		Did:      did,
		Points:   points,
		ToWhom:   toWhomProfile.Profile.DisplayNameNormalized,
		ToWhomID: toWhomProfile.ID,
		Because:  because,
	}, nil
}

func parseReasonAfterSpace(spaceNumber int, str string) string {
	for i := 0; i < spaceNumber; i++ {
		index := strings.Index(str, " ")
		str = str[index+1:]
	}
	return str
}
