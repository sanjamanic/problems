package problems

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if l > 1 && k > 0 {
		b := 0
		pos := b
		val := nums[pos]
		for cnt := 0; cnt < l; cnt++ {
			posNext := (pos + k) % l
			valNext := nums[posNext]
			nums[posNext] = val
			pos = posNext
			val = valNext
			if pos == b {
				pos++
				b = pos
				val = nums[pos]
			}
		}
	}
	return
}
