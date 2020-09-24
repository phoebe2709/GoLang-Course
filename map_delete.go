package main

import (
	"fmt"

)

func main() {
	m := map[string]int{
		"Phoebe": 32,
		"Sophie": 30,
	}
	fmt.Println(m)

	delete(m, "Phoebe")
	fmt.Println(m)

	delete(m, "Ian")
	fmt.Println(m)

	if v, ok := m["Sophie"]; ok {
		fmt.Println("value:", v)
		delete(m, "Sophie")
	}

	fmt.Println(m)
}
