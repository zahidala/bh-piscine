package piscine

// package main

// import "fmt"

func AlphaCount(s string) int {
	alphaRune := []rune(s)

	count := 0

	for _, code := range alphaRune {
		for i := 'a'; i <= 'z'; i++ {
			if i == code {
				count = count + 1
			}
		}

		for j := 'A'; j <= 'Z'; j++ {
			if j == code {
				count = count + 1
			}
		}
	}

	return count
}

// func main() {
// 	s := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
