// package piscine

package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	
	fmt.Println("a before division:", *a)
	fmt.Println("b before division:", *b)

	*a = *a / *b

	fmt.Println("a after division:", *a)
	fmt.Println("b after division:", *b)
	
	*b = *a % *b


}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
