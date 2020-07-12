func combine(n int, k int) [][]int {
	var ans [][]int
	if k == 0 {
		return ans
	}
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	backtrack(&ans, nums, 0, k)
	return ans
}

func backtrack(ans *[][]int, nums []int, first int, k int) {
	if first == k {
		*ans = append(*ans, append([]int{}, nums[0:k]...))
		return
	}
	for i := first; i < len(nums); i++ {
		if first == 0 || nums[i] > nums[first-1] {
			nums[i], nums[first] = nums[first], nums[i]
			backtrack(ans, nums, first+1, k)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	return
}