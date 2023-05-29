package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	toPrint := []rune{'x', ' ', '=', ' ', ('}' - 'I'), ('}' - 'K'), ',', ' ', 'y', ' ', '=', ' ', ('}' - 'K'), ('}' - 'L'), '\n'}
	for _, char := range toPrint {
		z01.PrintRune(char)
	}
}
