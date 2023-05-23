package piscine

// package main

// import "fmt"

func IsUpper(s string) bool {
	for i := range s {
		if s[i] <= 'A' || s[i] >= 'Z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsUpper("HELLO"))
// 	fmt.Println(IsUpper("HELLO!"))
// }
