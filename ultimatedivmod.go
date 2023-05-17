package piscine

// package main

// import "fmt"

func UltimateDivMod(a *int, b *int) {
	c := *a

	*a = c / *b

	*b = c % *b
}
