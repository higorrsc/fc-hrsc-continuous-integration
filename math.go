package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

// Soma returns the sum of two integers.
func Soma(a int, b int) int {
	return a + b
}
