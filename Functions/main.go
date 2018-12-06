package main

import "fmt"

func average(avarage_array []float64) float64 {
	var i int
	var result float64

	total := 0.0
	for i = 0; i < len(avarage_array); i++ {
		total = total + avarage_array[i]
	}

	result = total / float64(len(avarage_array))
	return result
}

func main() {
	first_array := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(first_array))
}
