package piscine

// package main 

// import "fmt"

func IsPrintable(s string) bool {
	for i := range s {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrintable("Hello"))
// 	fmt.Println(IsPrintable("Hello\n"))

// }
