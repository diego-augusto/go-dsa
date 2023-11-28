package godsa_test

import (
	godsa "go-dsa"
	"testing"
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
			r := godsa.IsAnagram_V2(tt.input_s, tt.input_t)
			if r != tt.output {
				t.Errorf("\ninput: %s %s\nexpected: %t\ngot: %t", tt.input_s, tt.input_t, tt.output, r)
			}
		})
	}
}
