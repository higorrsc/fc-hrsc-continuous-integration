package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
}

// Soma returns the sum of two integers.
func soma(a int, b int) int {
	return a + b
}
