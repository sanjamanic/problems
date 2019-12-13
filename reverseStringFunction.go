package problems

func reverseString(s []byte) {
	var c byte
	var p1, p2 *byte
	l := len(s)
	if l != 0 {
		for i := 0; i < l/2; i++ {
			p1 = &s[i]
			p2 = &s[l-i-1]
			c = *p1
			*p1 = *p2
			*p2 = c
		}
	}
	return
}
