// package main

// import "fmt"

package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	}
	if nb > 1 && nb < 21 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}

// func main() {
// 	arg := 0
// 	fmt.Println(RecursiveFactorial(arg))
// }
