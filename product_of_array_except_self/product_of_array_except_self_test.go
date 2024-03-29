package productofarrayexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductofArrayExceptSelf(t *testing.T) {
	testCases := []struct {
		desc   string
		intput []int
		output []int
	}{
		{
			desc:   "",
			intput: []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			desc:   "",
			intput: []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := productExceptSelf(tc.intput)
			assert.Equal(t, tc.output, r)
		})
	}
}
