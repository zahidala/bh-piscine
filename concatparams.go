package piscine

// package main

// import "fmt"

func ConcatParams(args []string) string {
	var words string
	for i := range args {
		if i == len(args)-1 {
			words = words + args[i]
		} else {
			words = words + args[i] + "\n"
		}
	}
	return words
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
