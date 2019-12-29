package main

import "testing"

type mockDatabase struct {
	val string
}

func (d *mockDatabase) Save(key string, value interface{}) error {
	d.val = key
	return nil
}

func TestParseCommand(t *testing.T) {
	type test struct {
		input         string
		expected      string
		expectedError string
	}
	tests := []test{
		{
			input:         "give kaden 250",
			expected:      "give 250",
			expectedError: "",
		},
		{
			input:         "give kaden",
			expected:      "",
			expectedError: "Give command expected 2 arguments, got 1",
		},
		{
			input:         "give kaden flajsldkf",
			expected:      "",
			expectedError: "Could not parse point value as integer",
		},
	}

	for _, test := range tests {
		storage := &mockDatabase{}
		server := &server{storage: storage}

		command, err := server.parseCommand(test.input)
		if len(test.expectedError) == 0 && err != nil {
			// wasn't expecting an error but got one anyway
			t.Error(err)
		} else if len(test.expectedError) > 0 && err.Error() != test.expectedError {
			// expecting an error but it was different from what was expected
			t.Errorf("\n expected error: %s\ngot: %s\n", test.expectedError, err.Error())
		}

		if err != nil {
			continue
		}

		command()

		if storage.val != test.expected {
			t.Errorf("\nexpected storage to have saved: %s\ngot: %s\n", test.expected, storage.val)
		}
	}
}
