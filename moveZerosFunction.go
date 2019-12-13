package problems

func moveZeroes(nums []int) {
	j := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < l; i++ {
		nums[i] = 0
	}
	return
}
