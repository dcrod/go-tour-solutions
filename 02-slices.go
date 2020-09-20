package main

import "golang.org/x/tour/pic"

func pattern(x, y, i int) uint8 {
	switch i {
	case 0:
		return uint8((x + y) / 2)
	case 1:
		return uint8(x * y)
	case 2:
		return uint8(x^(4*y)*x)
	case 3:
		return uint8(y^x * x^(3*y))
	default:
		return uint8(0)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = pattern(x, y, 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
