package main

import (
	"fmt"
)

func main() {
	a := (27 == 27)
	b := (27 <= 28)
	c := (27 >= 28)
	d := (27 != 28)
	e := (27 < 28)
	f := (27 > 28)
	fmt.Println(a, b, c, d, e, f)
}
