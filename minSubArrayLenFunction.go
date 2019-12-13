package problems

func minSubArrayLen(s int, nums []int) int {
	var sum []int
	var b, p2 int
	l := len(nums)
	if l > 0 {
		sum = make([]int, l)
		sum[0] = nums[0]
		sol := -1
		p1 := -1
		if nums[0] >= s {
			return 1
		}
		for i := 1; i < l; i++ {
			sum[i] = sum[i-1] + nums[i]
			if (sol == -1) && (sum[i] >= s) {
				sol = i + 1
				b = i
			}
			if nums[i] >= s {
				sol = 0
				break
			}
		}
		if sol > 0 {
			for p2 = b; p2 < l; p2++ {
				p1++
				for {
					if (sum[p2] - sum[p1]) >= s {
						sol = p2 - p1
						p1++
					} else {
						break
					}
				}
			}
			return sol
		} else {
			return sol + 1
		}
	} else {
		return 0
	}
}
