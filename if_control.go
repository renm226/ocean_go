package main

import "fmt"

func ifControlDemo() {
	x := 1
	if x == 1 {
		fmt.Println("x is equal to 1")
	} else {
		fmt.Println("x is not equal to 1")
	}
}

func main() {
	ifControlDemo()
}