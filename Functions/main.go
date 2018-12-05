package main

import "fmt"

func main() {
	first_array := []float64{98, 93, 77, 82, 83}
	var i int

	total := 0.0
	for i = 0; i < len(first_array); i++ {
		total = total + first_array[i]
	}

	total = total / float64(len(first_array))
	fmt.Println(total)
}
