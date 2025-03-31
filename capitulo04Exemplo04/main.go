package main

import (
	"fmt"
)

var x = 10
var y = 10.0

func main() {

	x = 20.5
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
}
