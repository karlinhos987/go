package main

import (
	"fmt"
)

func main() {
	x := 10
	y := "olá"

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", z, z)
}
