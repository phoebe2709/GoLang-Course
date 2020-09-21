package main

import (
	"fmt"

)

func main() {
	if true {
		fmt.Println("01")
	}
	if false {
		fmt.Println("02")
	}
	if !true {
		fmt.Println("03")
	}
	if !false {
		fmt.Println("04")
	}
	if 2 == 2 {
		fmt.Println("05")

	}
}
