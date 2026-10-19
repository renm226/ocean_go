package main

import "fmt"

func switchControlDemo() {

	x := 0

	switch x {
	case 1:
		fmt.Println("x is equal to 1")
	case 2:
		fmt.Println("x is equal to 2")
	default:
		fmt.Println("x is not equal to 1 or 2")
	}
}

func main() {
	switchControlDemo()
}