package main

import (
	"fmt"
	"math"
)

func swap_function(swap_x *int, swap_y *int) {
	*swap_x = 32
	*swap_y = 23
}

func distance(x1, y1, x2, y2 float64) float64 {
	var a float64
	var b float64

	a = x2 - x1
	b = y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var x, y int = 23, 32
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println("X = ", x, "\n", "Y = ", y)
	fmt.Println("\nSwap the numbers...Magic...\n", "Result: \n")
	swap_function(&x, &y)
	fmt.Println("X = ", x, "\n", "Y = ", y)

	fmt.Println("\n", rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
}
