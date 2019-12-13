package problems

func twoSum(numbers []int, target int) []int {
	var sol []int
	var i, j, m, n, val, jold int
	var flag bool
	sol = make([]int, 2)
	l := len(numbers)
	for i = 0; i < l; i++ {
		j = l - 1
		m = i + 1
		n = l - 1
		flag = false
		val = target - numbers[i]
		if val < numbers[i] {
			break
		}
		for {
			if val == numbers[j] {
				flag = true
				break
			} else if val < numbers[j] {
				jold = j
				n = j
				j = (j + m) / 2
			} else {
				jold = j
				m = j
				j = (j + n) / 2
			}
			if jold == j {
				break
			}
		}
		if flag == true {
			sol[0] = i + 1
			sol[1] = j + 1
			break
		}
	}
	return sol
}
