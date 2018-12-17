package main

import (
	"fmt"
	"math"
)

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

/*func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r
}*/

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	//var cx, cy, cr float64 = 0, 0, 5
	var c Circle

	fmt.Println("\n", rectangleArea(rx1, ry1, rx2, ry2))
	//fmt.Println(circleArea(cx, cy, cr))

	c.x = 10
	c.y = 5
	c.r = c.x * c.y
	fmt.Println("\nc.r = ", c.r)

	c = Circle{0, 0, 5}
	fmt.Println("\ncircleArea(c) = ", circleArea(c))
}
