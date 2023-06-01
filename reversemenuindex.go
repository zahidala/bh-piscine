// package piscine

package main

import "fmt"

func ReverseMenuIndex(menu []string) []string {
	reverseMenu := make([]string, len(menu))
	for i := 0; i <= len(menu)-1; i++ {
		reverseMenu[i] = menu[(len(menu)-1)-i]
	}
	return reverseMenu
}

func main() {
	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
