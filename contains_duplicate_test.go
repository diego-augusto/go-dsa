package godsa_test

import (
	godsa "go-dsa"
	"testing"
)

func TestContainsDuplicated(t *testing.T) {

	tc := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 3, 1},
			output: true,
		},
		{
			input:  []int{1, 2, 3, 4},
			output: false,
		},
		{
			input:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			output: true,
		},
		{
			input:  []int{1, 5, -2, -4, 0},
			output: false,
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {
			r := godsa.ContainsDuplicate_V2(tt.input)
			if r != tt.output {
				t.Errorf("\ninput: %d\nexpected: %t\ngot: %t", tt.input, tt.output, r)
			}
		})
	}
}
