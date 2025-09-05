package test

import (
	"testing"

	"github.com/matthieukhl/go-pokedex/utils"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    BoNjoUr MonDe ",
			expected: []string{"bonjour", "monde"},
		},
	}

	for _, c := range cases {
		actual := utils.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Got %d, want %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Got %s, want %s", word, expectedWord)
			}
		}
	}
}
