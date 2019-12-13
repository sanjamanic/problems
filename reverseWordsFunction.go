package problems

func reverseWords(s string) string {
	var res string
	var flag bool = false
	l := len(s)
	p1 := l - 1
	p2 := l - 1
	for {
		if flag {
			if len(res) > 0 {
				res = res + " "
			}
			res = res + s[p1+1:p2+1]
			p2 = p1
			flag = false
		}
		if p2 < 0 {
			break
		}
		for {
			if s[p1] == 32 && p1 > 0 {
				p1--
			} else {
				break
			}
		}
		p2 = p1

		for {
			if s[p1] != 32 && p1 > 0 {
				p1--
			} else {
				break
			}
		}
		if p1 == 0 && s[p1] != 32 {
			p1--
		}
		if p2 != p1 {
			flag = true
		} else {
			break
		}
	}
	return res
}
