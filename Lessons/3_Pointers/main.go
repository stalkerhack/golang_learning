package main

import "fmt"

func swap_function(swap_x *int, swap_y *int) {
	*swap_x = 32
	*swap_y = 23
}

func main() {
	var x, y int = 23, 32

	fmt.Println("X = ", x, "\n", "Y = ", y)
	fmt.Println("\nSwap the numbers...Magic...\n", "Result: \n")
	swap_function(&x, &y)
	fmt.Println("X = ", x, "\n", "Y = ", y)

}
