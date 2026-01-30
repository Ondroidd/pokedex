package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "lowercase and trim whitespace",
			input:    "  hElLo WORLd ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "multiple spaces between words",
			input:    "  hello    world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "single word",
			input:    "  HELLO  ",
			expected: []string{"hello"},
		},
		{
			name:     "empty string",
			input:    "   ",
			expected: []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanInput(tc.input)

			// Check length
			if len(actual) != len(tc.expected) {
				t.Errorf("Length mismatch: got %v words, want %v words",
					len(actual), len(tc.expected))
				t.Errorf("Got: %v", actual)
				t.Errorf("Want: %v", tc.expected)
				return
			}

			// Check each word
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("Word at index %d: got %q, want %q",
						i, actual[i], tc.expected[i])
				}
			}
		})
	}
}
