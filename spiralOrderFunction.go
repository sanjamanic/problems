package problems

func spiralOrder(matrix [][]int) []int {
	var i, h, l int
	var sol []int
	var flag int
	m := len(matrix)
	if m != 0 {
		n := len(matrix[0])
		k := m * n
		sol = make([]int, k)
		cnt := 0

		i = 0
		l = 0
		h = n - 1
		flag = 0
		for {
			sol, cnt = spiralLine(matrix, sol, i, l, h, flag, cnt)
			if cnt >= k {
				break
			}
			flag++
			if flag == 4 {
				flag = 0
			}
			switch flag {
			case 0:
				i = l
				h = n - 1
			case 1:
				i = h
				l = l + 1
				h = m - 1
				n--
			case 2:
				i = h
				l = l - 1
				h = n - 1
				m--
			case 3:
				i = l
				l = l + 1
				h = m - 1
			}

		}
		return sol
	} else {
		return nil
	}
}

func spiralLine(matrix [][]int, sol []int, i int, l int, h int, flag int, cnt int) ([]int, int) {
	var j int
	switch flag {
	case 0:
		for j = l; j <= h; j++ {
			sol[cnt] = matrix[i][j]
			cnt++
		}
	case 1:
		for j = l; j <= h; j++ {
			sol[cnt] = matrix[j][i]
			cnt++
		}
	case 2:
		for j = h; j >= l; j-- {
			sol[cnt] = matrix[i][j]
			cnt++
		}
	case 3:
		for j = h; j >= l; j-- {
			sol[cnt] = matrix[j][i]
			cnt++
		}
	}
	return sol, cnt
}
