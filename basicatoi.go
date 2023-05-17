package main

// package piscine

import "fmt"

func BasicAtoi(s string) int {
	numberArray := []rune(s)
	num := 0

	for i := 0; i < len(numberArray); i++ {
		num = num * 10 + (int(numberArray[i]) - 48)
	}
	return num
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
