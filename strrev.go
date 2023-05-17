package main

// package piscine

import "fmt"

func StrRev(s string) string {
	letterArray := []rune(s)

	var reversedLetters []rune

	for i := len(letterArray) - 1; i >= 0; i-- {
		reversedLetters = append(reversedLetters, letterArray[i])
	}

	return string(reversedLetters)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
