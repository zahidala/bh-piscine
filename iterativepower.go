// package main

package piscine

func IterativePower(nb int, power int) int {
	if nb == -1 {
		return -1
	} else if nb < 0 {
		return 0
	} else {
		result := 1
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
}

// func main() {
// 	fmt.Println(IterativePower(-1, -10))
// }
