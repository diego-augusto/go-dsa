package longestconsecutivesequence

func longestConsecutive(nums []int) int {

	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	longest := 0

	for _, n := range nums {
		if _, ok := m[n-1]; !ok {
			count := 1
			for {
				if _, ok := m[n+count]; !ok {
					break
				} else {
					count++
				}
			}
			if count > longest {
				longest = count
			}
		}
	}

	return longest
}
