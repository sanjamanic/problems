package problems

func pivotIndex(nums []int) int {
	var i, l, sumr int
	var flag bool = false
	l = len(nums)
	suml := make([]int, l+1)
	if l > 0 {
		suml[0] = 0
		for i = 1; i < l+1; i++ {
			suml[i] = suml[i-1] + nums[i-1]
		}
		for i = 0; i < l; i++ {
			sumr = suml[l] - suml[i+1]
			if suml[i] == sumr {
				flag = true
				break
			}
		}
	}
	if flag {
		return i
	} else {
		return -1
	}
}
