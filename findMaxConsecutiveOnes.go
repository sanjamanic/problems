package problems

func findMaxConsecutiveOnes(nums []int) int {
	k := 0
	sol := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			k = 0
		} else {
			k++
			if sol < k {
				sol = k
			}
		}
	}
	return sol
}
