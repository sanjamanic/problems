package problems

func findDiagonalOrder(matrix [][]int) []int {
	var i, j, k int
	var sol []int
	var flag bool = true
	M := len(matrix)
	if M != 0 {
		N := len(matrix[0])
		sol = make([]int, M*N)
		cnt := 0
		szMax := M
		if N < M {
		} else {
			szMax = N
		}
		for k = 0; k < szMax; k++ {
			for i = 0; i < k+1; i++ {
				j = k - i
				if j < 0 {
					break
				}
				switch flag {
				case true:
					if (j >= M) || (i >= N) {
						continue
					}
					sol[cnt] = matrix[j][i]
				case false:
					if (j >= N) || (i >= M) {
						continue
					}
					sol[cnt] = matrix[i][j]
				}
				cnt++
			}
			flag = !flag
		}

		for k = szMax; k <= 2*(szMax-1); k++ {
			for i = 1; i < szMax; i++ {
				j = k - i
				if j < 0 {
					continue
				}
				switch flag {
				case true:
					if (j >= M) || (i >= N) {
						continue
					}
					sol[cnt] = matrix[j][i]
				case false:
					if (j >= N) || (i >= M) {
						continue
					}
					sol[cnt] = matrix[i][j]
				}
				cnt++
			}
			flag = !flag
		}
		return sol
	} else {
		return nil
	}

}
