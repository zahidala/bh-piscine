package main

import "fmt"

// package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}
