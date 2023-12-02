package valid_anagram_test

import (
	"testing"

	"github.com/diego-augusto/go-dsa/valid_anagram"
)

func TestIsValidAnagram(t *testing.T) {

	tc := []struct {
		input_s string
		input_t string
		output  bool
	}{
		{
			input_s: "anagram",
			input_t: "nagaram",
			output:  true,
		},
		{
			input_s: "rat",
			input_t: "car",
			output:  false,
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {
			r := valid_anagram.IsAnagram_V2(tt.input_s, tt.input_t)
			if r != tt.output {
				t.Errorf("\ninput: %s %s\nexpected: %t\ngot: %t", tt.input_s, tt.input_t, tt.output, r)
			}
		})
	}
}
