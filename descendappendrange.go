// package piscine

package main

import "fmt"

func DescendAppendRange(max, min int) []int {
	var numArray []int

	if max <= min {
		numArray = []int{}
	}

	for i := max - 1; i > min-1; i-- {
		numArray = append(numArray, i+1)
	}
	return numArray
}

func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
	fmt.Println(DescendAppendRange(0, 1))
}
