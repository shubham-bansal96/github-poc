package main

import "fmt"

func main() {
	fmt.Println("github POC")
	fmt.Println("sum of two number is", add(5, 2))
}

func add(a, b int) int {
	return a + b
}
