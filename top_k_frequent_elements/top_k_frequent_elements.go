package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	arr := make([][]int, len(nums)+1)
	for kk, v := range m {
		arr[v] = append(arr[v], kk)
	}

	result := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 && k > 0 {
			result = append(result, arr[i]...)
			k = k - len(arr[i])
		}
	}

	return result
}
