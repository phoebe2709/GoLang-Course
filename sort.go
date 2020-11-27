package main

import (
	"fmt"
	"sort"

)

func main() {
	xi := []int{2, 5, 7, 8, 9, 4}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
}
