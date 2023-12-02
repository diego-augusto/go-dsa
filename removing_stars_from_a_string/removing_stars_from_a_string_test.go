package removing_stars_from_a_string_test

import (
	"testing"

	"github.com/diego-augusto/go-dsa/removing_stars_from_a_string"
	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "leet**cod*e",
			output: "lecoe",
		},
		{
			input:  "erase*****",
			output: "",
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			r := removing_stars_from_a_string.RemoveStars_V2(tc.input)
			assert.Equal(t, tc.output, r)
		})
	}
}
