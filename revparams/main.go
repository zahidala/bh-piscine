package main

import (
	"os"
	// "fmt"
	"github.com/01-edu/z01"
)

func main() {
	// arguments := os.Args
	// for i := 1; i <= len(arguments)-1; i++ {
	// 	fmt.Println((os.Args[i]))
	// }

	for i := len(os.Args) - 1; i >= 1; i-- {
		arguments := os.Args[i]
		for _, value := range arguments {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
