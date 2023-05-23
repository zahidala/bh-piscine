// package piscine

package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// reversedNumber := 0
	// for n != 0 {
	// 	reversedNumber = reversedNumber * 10 + n % 10
	// 	n = n / 10
	// }

	if n == 0 {
		z01.PrintRune('0')
	}

	for n != 0 {
		reversedNumber := rune(n%10 + '0') // 123 / 10 + '0'
		z01.PrintRune(reversedNumber)      // 3 + ascii 48
		n = n / 10                         // 3 + 48 = 51
		// break loop when n = 0			// rune(51) == the number 3
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
