package godsa

func IsAnagram_V1(s string, t string) bool {

	m := make(map[rune]int)

	for _, l := range s {
		m[l]++
	}

	for _, l := range t {
		m[l]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func IsAnagram_V2(s string, t string) bool {

	m := make([]int, 26)

	for _, l := range s {
		m[l-'a']++
	}

	for _, l := range t {
		m[l-'a']--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
