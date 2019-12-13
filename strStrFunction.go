package problems

func strStr(haystack string, needle string) int {
	var flag int = -1
	if len(needle) == 0 {
		return 0
	}
	j := 0
	if len(haystack) != 0 {
		for i := 0; i < len(haystack); {
			if haystack[i] == needle[j] {
				j++
				if j == len(needle) {
					flag = i - j + 1
					break
				}
			} else if j != 0 {
				i = i - j
				j = 0
			}
			i++
		}
	}
	return flag
}
