// package main

// import "fmt"

package piscine

func TrimAtoi(s string) int {
	numbers := 0
	isMinus := false

	for _, r := range s {
		if r == '-' && numbers == 0 && !isMinus {
			isMinus = true
		} else if r >= '0' && r <= '9' {
			numbers *= 10
			numbers += int(r - '0')
		}
	}
	if isMinus {
		return -numbers
	}
	return numbers
}

// func main() {
// 	fmt.Println(TrimAtoi("12345"))
// 	fmt.Println(TrimAtoi("str123ing45"))
// 	fmt.Println(TrimAtoi("012 345"))
// 	fmt.Println(TrimAtoi("Hello World!"))
// 	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
// }
