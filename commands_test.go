package main

import "testing"

type mockDatabase struct {
	change pointChange
}

func (d *mockDatabase) IncrementPoints(pc pointChange) error {
	d.change = pc
	return nil
}

func TestParseCommand(t *testing.T) {
	type test struct {
		command        pointCommand
		input          string
		expectedChange pointChange
		expectedError  string
	}
	tests := []test{
		{
			command:        GIVE,
			input:          "kaden 250 being a good boy",
			expectedChange: pointChange{User: "kaden", Points: 250, Reason: "being a good boy", UserChanging: "test"},
		},
		{
			command:       GIVE,
			input:         "kaden",
			expectedError: "Error: Too few arguments",
		},
		{
			command:       GIVE,
			input:         "kaden flajsldkf lfjasdklfjlakjsf",
			expectedError: "Could not parse point value as integer",
		},
		{
			command:       -1,
			input:         "kaden 250 being a good boy",
			expectedError: "unsupported command",
		},
	}

	for _, test := range tests {
		storage := &mockDatabase{}
		commandGenerator := newPointsCommandGenerator()
		command := commandGenerator(test.command, storage)

		pd := pointData{user: "test", arguments: test.input}
		err := command(pd)

		if len(test.expectedError) == 0 && err != nil {
			// wasn't expecting an error but got one anyway
			t.Error(err)
		} else if len(test.expectedError) > 0 && err.Error() != test.expectedError {
			// expecting an error but it was different from what was expected
			t.Errorf("\n expected error: %s\ngot: %s\n", test.expectedError, err.Error())
		}

		if storage.change != test.expectedChange {
			t.Errorf("\nexpected storage to have: %v\ngot: %v\n", test.expectedChange, storage.change)
		}
	}
}
