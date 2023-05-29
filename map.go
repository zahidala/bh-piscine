package piscine

// package main

func Map(f func(int) bool, a []int) []bool {
	var s []bool
	for _, v := range a {
		if f(v) {
			s = append(s, true)
		} else {
			s = append(s, false)
		}
	}
	return s
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	result := Map(piscine.IsPrime, a)
// 	fmt.Println(result)
// }
