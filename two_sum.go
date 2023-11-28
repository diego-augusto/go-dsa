package godsa

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{v, i}
		}
		m[n] = i
	}

	return []int{}
}
