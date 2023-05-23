// package piscine

package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// reversedNumber := 0
	// for n != 0 {
	// 	reversedNumber = reversedNumber * 10 + n % 10
	// 	n = n / 10
	// }

	var digits []int

	if n == 0 {
		z01.PrintRune('0')
	}

	for n != 0 {
		digit := n % 10
		digits = append(digits, digit)
		n = n / 10
	}

	// Print digits in ascending order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
	PrintNbrInOrder(4802622746086691776)
}
