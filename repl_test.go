package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur Pickachu",
			expected: []string{"charmander", "bulbasaur", "pickachu"},
		},
		{
			input:    "  Mew",
			expected: []string{"mew"},
		},
		{
			input:    "Evie      Magikarp      Snorlax",
			expected: []string{"evie", "magikarp", "snorlax"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("❌ expected a slice of length %d, but got a slice of %d for input %s", len(c.expected), len(actual), c.input)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected %s, got %s", expectedWord, word)
			}
		}
	}
}
