package main

import "testing"

type mockDatabase struct {
	user   string
	points int64
}

func (d *mockDatabase) IncrementPoints(user string, points int64) error {
	d.user = user
	d.points += points
	return nil
}

func TestParseCommand(t *testing.T) {
	type test struct {
		command        pointCommand
		input          string
		expectedUser   string
		expectedPoints int64
		expectedError  string
	}
	tests := []test{
		{
			command:        GIVE,
			input:          "kaden 250",
			expectedUser:   "kaden",
			expectedPoints: 250,
		},
		{
			command:       GIVE,
			input:         "kaden",
			expectedError: "Give command expected 2 arguments, got 1",
		},
		{
			command:       GIVE,
			input:         "kaden flajsldkf",
			expectedError: "Could not parse point value as integer",
		},
		{
			command:       -1,
			input:         "kaden 250",
			expectedError: "unsupported command",
		},
	}

	for _, test := range tests {
		storage := &mockDatabase{}
		// server := &server{storage: storage}
		commandGenerator := newPointsCommandGenerator()
		command := commandGenerator(test.command, storage)

		err := command(test.input)

		if len(test.expectedError) == 0 && err != nil {
			// wasn't expecting an error but got one anyway
			t.Error(err)
		} else if len(test.expectedError) > 0 && err.Error() != test.expectedError {
			// expecting an error but it was different from what was expected
			t.Errorf("\n expected error: %s\ngot: %s\n", test.expectedError, err.Error())
		}

		if storage.user != test.expectedUser {
			t.Errorf("\nexpected storage to have updated user: %s\ngot: %s\n", test.expectedUser, storage.user)
		}
		if storage.points != test.expectedPoints {
			t.Errorf("\nexpected storage to have updated points: %d\ngot: %d\n", test.expectedPoints, storage.points)
		}
	}
}
