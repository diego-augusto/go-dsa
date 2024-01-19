package groupshiftedstrings

import (
	"strconv"
	"strings"
)

func getShiftedHash(w string) string {

	if len(w) < 2 {
		return ""
	}

	output := make([]string, 0)

	for i := 1; i < len(w); i++ {

		first := rune(w[i-1])
		second := rune(w[i])

		shift := int(first - second)

		if shift < 0 {
			shift = shift * -1
		}

		if first > second {
			r := int(26 - shift)
			output = append(output, strconv.Itoa(r))
		} else {
			output = append(output, strconv.Itoa(shift))
		}

	}

	return strings.Join(output, "")

}

func groupShiftedStrings(input []string) [][]string {

	m := make(map[string][]string)

	for _, s := range input {
		shift := getShiftedHash(s)
		m[shift] = append(m[shift], s)
	}

	output := make([][]string, 0)

	for _, v := range m {

		output = append(output, v)

	}

	return output
}
