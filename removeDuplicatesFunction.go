package problems

func removeDuplicates(nums []int) int {
	j := 1
	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
