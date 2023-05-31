package piscine

// package main

func Unmatch(a []int) int {
	var compareNum int
	for i := 0; i < len(a); i++ {
		if compareNum < a[i] {
			compareNum = a[i]
		}
	}
	return compareNum
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
