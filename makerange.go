package piscine

// package main

// import "fmt"

func MakeRange(min, max int) []int {
	var numArray []int

	if max > min {
		numArray = make([]int, max-min)
		for i := range numArray {
			numArray[i] = i + min
		}
	}
	return numArray
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
