package problems

func dominantIndex(nums []int) int {
	var max1, max2, sol, i int = 0, 0, 0, 0

	for i = 0; i < len(nums); i++ {
		if max2 < nums[i] {
			max1 = max2
			max2 = nums[i]
			sol = i
		} else if max1 < nums[i] {
			max1 = nums[i]
		}
	}
	if max2 >= 2*max1 {
		return sol
	} else {
		return -1
	}
}
