package piscine

// package main

// func PrintNbr(nbr int) {
// 	z01.PrintRune(rune(nbr) + 48)
// }

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	ForEach(PrintNbr, a)
// }
