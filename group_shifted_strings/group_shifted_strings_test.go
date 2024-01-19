package groupshiftedstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {

	testCases := []struct {
		desc   string
		input  string
		output string
	}{
		{
			desc:   "asc sequence",
			input:  "ab",
			output: "1",
		},
		{
			desc:   "desc sequence",
			input:  "ba",
			output: "25",
		},
		{
			desc:   "asc sequence 2",
			input:  "az",
			output: "25",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := getShiftedHash(tc.input)
			assert.Equal(t, tc.output, n)
		})
	}

}

func TestGroupShiftedStrings(t *testing.T) {

	testCases := []struct {
		desc   string
		input  []string
		output [][]string
	}{
		{
			desc:   "sequence",
			input:  []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			output: [][]string{{"abc", "bcd", "xyz"}, {"az", "ba"}, {"acef"}, {"a", "z"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := groupShiftedStrings(tc.input)
			assert.ObjectsAreEqualValues(tc.output, n)
		})
	}

}
