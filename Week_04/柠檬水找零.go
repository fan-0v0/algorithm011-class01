func lemonadeChange(bills []int) bool {
	billMap := map[int]int{0: -1, 5: 0, 10: 0, 20: 0}
	for _, val := range bills {
		change := val - 5
		if change == 5 {
			if billMap[change] == 0 {
				return false
			}
			billMap[change]--
		} else if change == 15 {
			if billMap[10] == 0 {
				if billMap[5] < 3 {
					return false
				}
				billMap[5] = billMap[5] - 3
			} else if billMap[5] == 0 {
				return false
			} else {
				billMap[5]--
				billMap[10]--
			}
		}
		billMap[val]++
	}
	return true
}