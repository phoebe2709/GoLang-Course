package main

import (
	"fmt"
)

func main() {
	favSport := "combat"

	switch favSport {

	case "horseriding":
		fmt.Println("neehhh")
	case "dancing":
		fmt.Println("night fever")
	case "swimming":
		fmt.Println("go to the pool")
	case "combat":
		fmt.Println("fight!!!")
	}
}
