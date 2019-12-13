package problems

func addBinary(a string, b string) string {
	var sol, c string
	var prev int
	la := len(a)
	lb := len(b)
	if la > lb {
		c = a
		lc := la
		a = b
		b = c
		la = lb
		lb = lc
	}
	prev = 0
	for i := 0; i < la; i++ {
		switch a[la-1-i] {
		case 48:
			switch b[lb-1-i] {
			case 48:
				switch prev {
				case 0:
					sol = "0" + sol
					prev = 0
				case 1:
					sol = "1" + sol
					prev = 0
				}
			case 49:
				switch prev {
				case 0:
					sol = "1" + sol
					prev = 0
				case 1:
					sol = "0" + sol
					prev = 1
				}
			}
		case 49:
			switch b[lb-1-i] {
			case 48:
				switch prev {
				case 0:
					sol = "1" + sol
					prev = 0
				case 1:
					sol = "0" + sol
					prev = 1
				}
			case 49:
				switch prev {
				case 0:
					sol = "0" + sol
					prev = 1
				case 1:
					sol = "1" + sol
					prev = 1
				}
			}
		}

	}
	for i := la; i < lb; i++ {
		switch b[lb-1-i] {
		case 48:
			switch prev {
			case 0:
				sol = "0" + sol
				prev = 0
			case 1:
				sol = "1" + sol
				prev = 0
			}
		case 49:
			switch prev {
			case 0:
				sol = "1" + sol
				prev = 0
			case 1:
				sol = "0" + sol
				prev = 1
			}
		}
	}
	if prev == 1 {
		sol = "1" + sol
	}
	return sol

}
