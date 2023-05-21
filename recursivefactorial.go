// package main

package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb > 1 && nb < 65535 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}

// func main() {
// 	arg := 4
// 	fmt.Println(RecursiveFactorial(arg))
// }
