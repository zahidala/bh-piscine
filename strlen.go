package main
// package piscine

import "fmt"

func StrLen(s string) int {
	
	// count := 0

	// for range s {
	// 	count++
	// }
	// return count

	s2 := []rune(s)
	return len(s2)
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
