package godsa_test

import (
	godsa "go-dsa"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {

	tc := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {
			r := godsa.IsValidParentheses(tt.input)
			if r != tt.output {
				t.Errorf("\ninput: %s\nexpected: %t\ngot: %t", tt.input, tt.output, r)
			}
		})
	}
}
