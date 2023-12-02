package valid_parentheses_test

import (
	"testing"

	"github.com/diego-augusto/go-dsa/valid_parentheses"
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
			r := valid_parentheses.IsValidParentheses(tt.input)
			if r != tt.output {
				t.Errorf("\ninput: %s\nexpected: %t\ngot: %t", tt.input, tt.output, r)
			}
		})
	}
}
