package remove_duplicates_test

import (
	"testing"

	"github.com/diego-augusto/go-dsa/remove_duplicates"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output []int
	}{
		{
			input:  []int{0, 0, 0, 1, 1},
			output: []int{0, 1},
		},
		{
			input:  []int{1, 1, 2},
			output: []int{1, 2},
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := remove_duplicates.RemoveDuplicates(tc.input)
			assert.Equal(t, tc.output, tc.input[:n])
		})
	}
}
