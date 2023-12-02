package remove_duplicates

func RemoveDuplicates(nums []int) int {

	p := 0

	for i := 1; i < len(nums); i++ {
		if nums[p] != nums[i] {
			p++
			nums[p] = nums[i]
		}
	}

	return p + 1
}
