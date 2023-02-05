package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	z01.PrintRune(52)
	z01.PrintRune(50)

	printStr(", y = ")
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune('\n')
}
