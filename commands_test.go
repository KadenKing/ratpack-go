package main

import "testing"

import "errors"

type fakeTranslater struct{}

func (p *fakeTranslater) GetProfileByID(id string) (string, error) {
	if id == "1234" {
		return "tester", nil
	}
	return "", errors.New("could not find user with that id")
}
func (p *fakeTranslater) GetProfileByUsername(username string) (slackMembersResponse, error) {
	if username == "kaden" {
		return slackMembersResponse{ID: "1234", Profile: slackProfile{DisplayNameNormalized: "thrifty watermelon"}, Name: "kaden.king.king"}, nil
	}
	return slackMembersResponse{}, errors.New("could not find user")
}

func TestParseCommand(t *testing.T) {

	type test struct {
		input              slackRequest
		expectedWhoDidWhat whoDidWhat
		expectedError      string
	}
	tests := []test{
		{
			input:              slackRequest{Text: "kaden 250 being good", UserID: "1234"},
			expectedWhoDidWhat: whoDidWhat{Who: "tester", WhoID: "1234", Did: "gave", Points: 250, ToWhom: "thrifty watermelon", Because: "being good", ToWhomID: "1234"},
		},
		{
			input:         slackRequest{Text: "kaden"},
			expectedError: "Error: Too few arguments",
		},
		{
			input:         slackRequest{Text: "kaden fjlaksjflk being good"},
			expectedError: "Could not parse point value as integer",
		},
		{
			input:         slackRequest{Text: "test 250 being bad", UserID: "1234"},
			expectedError: "could not find user",
		},
		{
			input:         slackRequest{Text: "test 250 being bad", UserID: "4567"},
			expectedError: "could not find user with that id",
		},
	}

	for _, test := range tests {
		parser := giveCommandParser{}
		wdw, err := parser.Parse(test.input, &fakeTranslater{})

		if len(test.expectedError) == 0 && err != nil {
			// wasn't expecting an error but got one anyway
			t.Error(err)
		} else if len(test.expectedError) > 0 && err.Error() != test.expectedError {
			// expecting an error but it was different from what was expected
			t.Errorf("\n expected error: %s\ngot: %s\n", test.expectedError, err.Error())
		}

		if wdw != test.expectedWhoDidWhat {
			t.Errorf("\nexpected who did what to be: %v\ngot: %v\n", test.expectedWhoDidWhat, wdw)
		}
	}
}
