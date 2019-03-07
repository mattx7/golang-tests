package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n==> Sets")
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["Ann"] { // will be false if person is not in the map
		fmt.Println("Ann", "was at the meeting")
	}

	fmt.Println("Map content:")
	for key, value := range attended {
		fmt.Printf("key: %s, value: %t\n", key, value)
	}
}
