package main

import (
	"fmt"

)

func main() {
	x := "whos mad"
	if x == "Phoebe Rowland" {
		fmt.Println(x)
	} else if x == "Sophie Rowland" {
		fmt.Println("Sophie is mad")
	} else {
		fmt.Println("Were both mad")
	}
}
