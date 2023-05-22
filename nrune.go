package piscine

// package main

func NRune(s string, n int) rune {
	r := []rune(s)

	if n == -1 {
		return r[len(s)-1]
	} else if n < len(s) || n == len(s)-1 {
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
