package piscine

// package main

// import "fmt"

func IsNumeric(s string) bool {
	for i := range s {
		if (s[i] < '0' || s[i] > '9') {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsNumeric("010203"))
// 	fmt.Println(IsNumeric("01,02,03"))
// }
