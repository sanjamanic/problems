package problems

func longestCommonPrefix(strs []string) string {
	var h, c string
	var i, j, lc int = 0, -1, 0
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}
	c = strs[0]
	lc = len(c)
	if lc == 0 {
		return ""
	}
	for {
		for i = 1; i < l; i++ {
			if lc < j+2 {
				break
			}
			h = strs[i]
			if len(h) < j+2 {
				break
			}
			if h[j+1] != c[j+1] {
				break
			}
		}
		if i == l {
			j++
		} else {
			break
		}
	}
	if j < 0 {
		return ""
	} else {
		return c[0 : j+1]
	}
}
