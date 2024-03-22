package top_k_frequent_elements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		k    int
		out  []int
	}{
		{
			in:  []int{1, 1, 1, 2, 2, 3},
			k:   2,
			out: []int{1, 2},
		},
		{
			in:  []int{1},
			k:   1,
			out: []int{1},
		},
		{
			in:  []int{1, 2},
			k:   2,
			out: []int{1, 2},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			r := topKFrequent(tc.in, tc.k)
			assert.ElementsMatch(t, tc.out, r)
		})
	}
}
