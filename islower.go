package piscine

// package main

// import "fmt"

func IsLower(s string) bool {
	for i := range s {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsLower("hello"))
// 	fmt.Println(IsLower("hello!"))
// }
