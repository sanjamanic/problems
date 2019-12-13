package problems

func arrayPairSum(nums []int) int {
	n := len(nums) / 2
	sum := 0
	numsSort := sortArray(nums)
	for i := 0; i < n; i++ {
		sum = sum + numsSort[i]
	}
	return sum
}
func sortArray(nums []int) []int {
	var numsSort []int
	var k, p int = 0, 0
	var minVal, minValp int
	n := len(nums)
	numsSort = make([]int, n/2)

	for i := 0; i < n/2; i++ {
		minVal = nums[0]
		minValp = nums[1]
		k = 0
		p = 1
		if minValp < minVal {
			minVal = nums[1]
			minValp = nums[0]
			k = 1
			p = 0
		}
		for j := 2; j < n; j++ {
			if nums[j] < minVal {
				minValp = minVal
				minVal = nums[j]
				p = k
				k = j
			} else if nums[j] <= minValp {
				minValp = nums[j]
				p = j
			}
		}
		numsSort[i] = minVal
		nums[k] = 10001
		nums[p] = 10001
	}
	return numsSort
}
