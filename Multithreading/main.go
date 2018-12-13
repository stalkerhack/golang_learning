package main

import (
	"fmt"
	"math/rand"
	"time"
)

func first__function_gorutine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func second_function_gorutine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	go first_gorutine(123)
	fmt.Println("\n")
	//////////////////////////////////////////
	for i := 0; i < 10; i++ {
		go second_function_gorutine(i)
	}

	var input string
	fmt.Scanln(&input)
}
