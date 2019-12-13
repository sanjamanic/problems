package problems

func getRow(rowIndex int) []int {
	var triangle []int

	triangle = make([]int, rowIndex+1)
	triangle[0] = 1
	if rowIndex > 0 {
		triangle[1] = 1
		for i := 2; i <= rowIndex; i++ {
			triangle[i] = 1
			for j := i - 1; j > 0; j-- {
				triangle[j] = triangle[j-1] + triangle[j]
			}
		}
	}
	return triangle
}
