package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello     world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloWorld",
			expected: []string{"helloworld"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "Hello World How Are You",
			expected: []string{"hello", "world", "how", "are", "you"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of actual: %v does not match expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual: %v does not match expected: %v", word, expectedWord)
			}
		}
	}
}
