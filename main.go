package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println("Second Branch")
	fmt.Println(calc(1, 5))
}

func calc(x, y int) int {
	return x + y
}
