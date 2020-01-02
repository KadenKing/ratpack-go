package main

import "testing"

type fakeTranslater struct{}

func (p *fakeTranslater) GetProfileByID(id string) (string, error) {
	return "tester", nil
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
			expectedWhoDidWhat: whoDidWhat{who: "tester", did: "gave", points: 250, toWhom: "kaden", because: "being good"},
		},
		{
			input:         slackRequest{Text: "kaden"},
			expectedError: "Error: Too few arguments",
		},
		{
			input:         slackRequest{Text: "kaden fjlaksjflk being good"},
			expectedError: "Could not parse point value as integer",
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
