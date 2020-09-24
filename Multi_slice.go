package main

import (
	"fmt"
)

func main() {
	jb := []string{"Phoebe", "one", "two"}
	fmt.Println(jb)
	mp := []string{"Rowland", "three", "four"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
