package problems

func plusOne(digits []int) []int {
	var sol []int
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	if digits[0] == 0 {
		sol = make([]int, l+1)
		sol[0] = 1
		for i := 0; i < l; i++ {
			sol[i+1] = digits[i]
		}
	} else {

		sol = digits
	}

	return sol
}
