package godsa_test

import (
	godsa "go-dsa"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tc := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {
			r := godsa.TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(r, tt.output) {
				t.Errorf("\ninput: %d %d\nexpected: %d\ngot: %d", tt.nums, tt.target, tt.output, r)
			}
		})
	}
}
