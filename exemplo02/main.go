package main

import (
	"fmt"
)

// variaveis que podem ser utilizadas em qualquer parte do codigo.
var y = 10

func main() {
	z := 20
	qualquerCoisa(z)
}

func qualquerCoisa(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
