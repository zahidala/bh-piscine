package main

// package piscine

import "fmt"

func IterativePower(nb int, power int) int {
	if nb == -1 {
		return -1
	} else {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
			if result < 0 {
				return 0
			}
		}
		return result
	}
}

func main() {
	fmt.Println(IterativePower(-1, 4))
}
