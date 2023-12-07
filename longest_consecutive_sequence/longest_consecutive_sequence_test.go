package longestconsecutivesequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			input:  []int{100, 4, 200, 1, 3, 2},
			output: 4,
		},

		{
			input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			output: 9,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := longestConsecutive(tc.input)
			assert.Equal(t, tc.output, r)
		})
	}
}
