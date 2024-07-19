package answers

func jump(nums []int) int {
	l, r := 0, 0
	steps := 0

	for r < len(nums)-1 {
		farthest := 0

		for i := l; i < r+1; i++ {
			farthest = max(farthest, i+nums[i])
		}

		l = r + 1
		r = farthest
		steps++
	}
	return steps
}
