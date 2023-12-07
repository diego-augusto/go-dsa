package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	arr := make([][]int, len(nums)+1)
	for key, value := range m {
		arr[value] = append(arr[value], key)
	}

	r := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		r = append(r, arr[i]...)
		if len(r) == k {
			break
		}
	}

	return r
}
