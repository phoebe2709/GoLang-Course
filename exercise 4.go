package main

import (
	"fmt"

)

type Phoebe int

var x Phoebe

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
