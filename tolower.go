package piscine

// package main

// import "fmt"

func ToLower(s string) string {
	x := []rune(s)
	for i := range x {
		if x[i] >= 'A' && x[i] <= 'Z' {
			x[i] = x[i] + 32
		}
	}
	return string(x)
}

// func main() {
// 	fmt.Println(ToLower("Hello! How are you?"))
// }
