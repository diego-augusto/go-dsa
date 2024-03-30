package stringcompression

import "fmt"

func compress(chars []byte) int {

	c := chars[0]
	count := 1
	p := 0

	if len(chars) == 1 {
		return 1
	}

	for i := 1; i < len(chars); i++ {
		if chars[i] == c {
			count++
		} else {
			if count > 1 {
				chars[p] = c
				p++
				for _, b := range fmt.Sprint(count) {
					chars[p] = byte(b)
					p++
				}
			} else {
				chars[p] = c
				p++
			}
			count = 1
			c = chars[i]
		}
	}

	if count > 1 {
		chars[p] = c
		p++
		for _, b := range fmt.Sprint(count) {
			chars[p] = byte(b)
			p++
		}
	} else {
		chars[p] = c
		p++
	}

	return p

}
