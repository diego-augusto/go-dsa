package productofarrayexceptself

func productExceptSelf(nums []int) []int {

	ans := make([]int, len(nums))
	p := 1
	for i, n := range nums {
		ans[i] = p
		p *= n
	}

	p = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= p
		p *= nums[i]
	}

	return ans
}
