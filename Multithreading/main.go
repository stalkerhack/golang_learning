package main

import "fmt"

func first_gorutine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go first_gorutine(123)
	var input byte
	fmt.Scanln(&input)
}
