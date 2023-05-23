package piscine

// package main

// import "fmt"

func ToUpper(s string) string {
	x := []rune(s)
	for i := range x {
		if x[i] >= 'a' && x[i] <= 'z' {
			x[i] = x[i] - 32
		}
	}
	return string(x)
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }