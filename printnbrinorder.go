package piscine

// package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// reversedNumber := 0
	// for n != 0 {
	// 	reversedNumber = reversedNumber * 10 + n % 10
	// 	n = n / 10
	// }

	// if n == 0 {
	// 	z01.PrintRune('0')
	// }

	// for n != 0 {
	// 	reversedNumber := rune(n%10 + '0') // 123 / 10 + '0'
	// 	z01.PrintRune(reversedNumber)      // 3 + ascii 48
	// 	n = n / 10                         // 3 + 48 = 51
	// 	// break loop when n = 0			// rune(51) == the number 3
	// }

	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var array []int
		eachValue := 0

		arrayCount := 0
		var minValue int
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		for count := range array {
			arrayCount = count + 1
		}
		for i := 0; i < arrayCount-1; i++ {
			for j := 0; j < arrayCount-i-1; j++ {
				if array[j] > array[j+1] {
					minValue = array[j]
					array[j] = array[j+1]
					array[j+1] = minValue
				}
			}
		}
		for i := 0; i < arrayCount; i++ {
			z01.PrintRune(rune(array[i] + 48))
		}
	}
}

// func main() {
// 	PrintNbrInOrder(321)
// 	PrintNbrInOrder(0)
// 	PrintNbrInOrder(321)
// 	PrintNbrInOrder(4802622746086691776)
// }
