package piscine

// package main

// import "fmt"

func Index(s string, toFind string) int {
	lettersLen := len([]rune(s))
	letterLen := len([]rune(toFind))

	for i := 0; i <= lettersLen-letterLen; i++ {
		if s[i:i+letterLen] == toFind {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
