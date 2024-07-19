package answers

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int) // val -> index

	for i, n := range nums {
		diff := target - n
		if _, ok := prevMap[diff]; ok {
			return []int{prevMap[diff], i}
		}
		prevMap[n] = i
	}

	return []int{}
}
