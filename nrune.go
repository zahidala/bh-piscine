package piscine

// package main

func NRune(s string, n int) rune {
	r := []rune(s)

	count := 0

	for i := range s {
		count = i + 1
	}

	if n <= count && n-1 >= 0 {
		return r[n-1]
	} else {
		return 0
	}
}

// func main() {
// 	z01.PrintRune(NRune("Hello!", 3))
// 	z01.PrintRune(NRune("Salut!", 2))
// 	z01.PrintRune(NRune("Bye!", -1))
// 	z01.PrintRune(NRune("Bye!", 5))
// 	z01.PrintRune(NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }
