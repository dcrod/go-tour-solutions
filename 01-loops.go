// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	i := 0
	z, diff := x/4, 1.0
	for diff * diff > 1e-12 {
		diff = (z*z - x) / (2*z)
		z -= diff
		i++
	}
	fmt.Println("Took", i, "loops")
	return z
}

func main() {
	val := float64(7493759313)
	fmt.Println(Sqrt(val))
	fmt.Println(math.Sqrt(val))
}
