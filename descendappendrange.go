package piscine

// package main

func DescendAppendRange(max, min int) []int {
	var numArray []int

	for i := max - 1; i > min-1; i-- {
		numArray = append(numArray, i+1)
	}
	return numArray
}

// func main() {
// 	fmt.Println(DescendAppendRange(10, 5))
// 	fmt.Println(DescendAppendRange(5, 10))
// }
