package main

import (
	"fmt"
)

func main() {
	s := "HELLO WORLD"
	fmt.Printf(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}
