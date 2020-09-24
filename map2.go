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

	fmt.Println(m["Phoebe"])

	fmt.Println(m["Sue"])

	v, ok := m["Sue"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Todd"] = 33

	if v, ok := m["Sue"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
