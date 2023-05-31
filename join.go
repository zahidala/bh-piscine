package piscine

// package main

// import "fmt"

func Join(strs []string, sep string) string {
	words := strs[0]
	for i := 1; i < len(strs); i++ {
		words += sep + strs[i]
	}
	return words
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
