func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if _, ok := m[val]; ok {
			ans := []int{m[val], i}
			return ans
		} else {
			m[target-val] = i
		}
	}
	return nil
}