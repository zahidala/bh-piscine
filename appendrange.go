package piscine

// package main

import "fmt"

func AppendRange(min, max int) []int {
	var numArray []int

	for i := min - 1; i < max-1; i++ {
		numArray = append(numArray, i+1)
	}
	return numArray
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 5))
// }
