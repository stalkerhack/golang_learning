package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var today time.Weekday

	fmt.Println("When is Sunday?")
	today = time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.\n")
	case today + 1:
		fmt.Println("Tommorow.\n")
	case today + 2:
		fmt.Println("In two days.\n")
	default:
		fmt.Println("Too far away.\n")
	}
	fmt.Println(reflect.TypeOf(today))
}
