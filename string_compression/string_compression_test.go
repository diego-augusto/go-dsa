package stringcompression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCompression(t *testing.T) {
	testCases := []struct {
		desc   string
		intput []byte
		output int
	}{
		{
			desc:   "",
			intput: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			output: 6,
		},
		{
			desc:   "",
			intput: []byte{'a'},
			output: 1,
		},
		{
			desc:   "",
			intput: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			output: 4,
		},
		{
			desc:   "",
			intput: []byte{'a', 'b', 'c'},
			output: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := compress(tc.intput)
			assert.Equal(t, tc.output, r)
		})
	}
}
