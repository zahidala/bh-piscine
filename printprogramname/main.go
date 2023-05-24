package main

import (
	"os"
	// "fmt"
	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, value := range name {
		if value != '.' && value != '/' {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}
