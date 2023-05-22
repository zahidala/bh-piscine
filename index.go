// package piscine

package main

import "fmt"

func Index(s string, toFind string) int {
	for i, letters := range s {
		for _, letter := range toFind {
			if letters == letter {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
