package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	greater := false
	lower := false
	equal := false
	if len(a) == 1 {
		return true
	}
	for i := 0; i < len(a)-1; i++ {
		com := f(a[i], a[i+1])
		if com > 0 {
			greater = true
		} else if com < 0 {
			lower = true
		} else if com == 0 {
			equal = true
		}
	}
	if greater != lower {
		return true
	} else if greater == lower && greater != false {
		return false
	}
	return equal

}
