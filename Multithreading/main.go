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
func message(c chan string) {
	for i := 0; ; i++ {
		c <- "Message..."
	}
}
func message_output(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)

	go first__function_gorutine(123)
	fmt.Println("\n")
	//////////////////////////////////////////
	for i := 0; i < 10; i++ {
		go second_function_gorutine(i)
	}
	//////////////////////////////////////////
	go message(c)
	go message_output(c)

	var input string
	fmt.Scanln(&input)
}
