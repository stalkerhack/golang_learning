package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
	fmt.Println("*xPtr = ", *xPtr)
	fmt.Println("xPtr = ", xPtr)
	fmt.Println("&xPtr = ", &xPtr)
}
func main() {
	x := 5
	zero(&x)
	fmt.Println("\nx = ", x)
	fmt.Println("&x = ", &x)
}
