package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1:]

	for i := 0; i <= len(s)-1; i++ {
		if s[i] == "01" {
			fmt.Println("Alert!!!")
		}

		if s[i] == "01" && s[i] == "galaxy" {
			fmt.Println("Alert")
		}

		if s[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
