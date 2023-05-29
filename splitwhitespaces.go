package piscine

// package main

// import "fmt"

func SplitWhiteSpaces(s string) []string {
	var letter string
	var words []string

	for i := range s {
		// fmt.Println(string(s[i]))
		if s[i] != 32 && s[i] != 9 && s[i] != 10 {
			letter = letter + string(s[i])
		}

		if (s[i] == 32 || s[i] == 9 || s[i] == 10 || i == len(s)-1) && len(letter) > 0 {
			words = append(words, letter)
			letter = ""
		}
	}
	return words
}

// func main() {
// 	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
// }
