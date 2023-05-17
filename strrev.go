// package main

package piscine

// import "fmt"

func StrRev(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]

	}

	return string(rns)
}

// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
