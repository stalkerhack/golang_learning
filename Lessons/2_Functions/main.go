package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////
// Объявление и описание функции avarage
func average(avarage_array []float64) (result float64) {
	var i int
	// Сложение всех элементов массива
	total := 0.0
	for i = 0; i < len(avarage_array); i++ {
		total = total + avarage_array[i]
	}
	// Расчет среднего арифметического и запись в переменную result
	result = total / float64(len(avarage_array))
	return
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// Объявление и описание функции factorial
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// Объявление и описание функции fibonacci.....Доделать!!!!!!
func fibonacci(input int) int {

	if input == 0 {
		return input
	}
	if input == 1 {
		return input
	}

	return fibonacci(input-1) + fibonacci(input-2)
}

////////////////////////////////////////////////////////////////////////////////
func main() {
	var i int
	var loop int
	var input_value_fibonacci int
	var input_value_factorial uint

	fmt.Println("Введите число для расчета факториала: ")
	fmt.Scanln(&input_value_factorial)

	fmt.Println("\nВведите число для числе Фибоначчи: ")
	fmt.Scanln(&input_value_fibonacci)

	// Объявление массива, который будет передан функции avarage
	first_array := []float64{98, 93, 77, 82, 83}

	// Вызов функции avarage, считающей среднее арифметическое
	fmt.Println("\n\nСреднее арифметическое = ", average(first_array))

	// Вызов функции factorial, считающей факториал
	fmt.Println("Факториал числа = ", factorial(input_value_factorial))

	//Вызов функции fibonacci, считающей последовательность числе Фибоначчи
	fmt.Println("Числа Фибоначчи: ")
	for loop = 1; loop < input_value_fibonacci; loop++ {
		fmt.Println(fibonacci(i))
		i++
	}
}
