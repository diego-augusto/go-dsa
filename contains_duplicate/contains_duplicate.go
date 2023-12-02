package contains_duplicate

func ContainsDuplicate_V1(nums []int) bool {

	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}

func ContainsDuplicate_V2(nums []int) bool {
	m := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}

	return false
}
