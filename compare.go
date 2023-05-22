// package main

package piscine

func Compare(a, b string) int {
	// for _, i := range a {
	// 	for _, j := range b {
	// 		if i == j {
	// 			return 0
	// 		} else if i > j {
	// 			return 1
	// 		} else {
	// 			return -1
	// 		}
	// 	}
	// }
	// return 0

	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

// func main() {
// 	fmt.Println(Compare("Hello!", "Hello!"))
// 	fmt.Println(Compare("Salut!", "lut!"))
// 	fmt.Println(Compare("Ola!", "Ol"))
// }
