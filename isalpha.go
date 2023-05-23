package main

// package piscine

import "fmt"

func IsAlpha(s string) bool {

	for i := range s {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && (s[i] < 0 || s[i] > 9){
			
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
