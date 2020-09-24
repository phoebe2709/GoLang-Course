package main

import (
	"fmt"

)

func main() {
	x := []int{4, 5, 7, 8, 27}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{272, 234, 222, 833}
	x = append(x, y...)
	fmt.Print(x)
}
