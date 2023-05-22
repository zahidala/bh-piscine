package piscine

// package main

import "unicode/utf8"

func AlphaCount(s string) int {
	return utf8.RuneCountInString(s)
}

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
