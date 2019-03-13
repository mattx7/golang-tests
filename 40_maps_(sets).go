package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n==> Sets")
	attendedSet := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attendedSet["Ann"] { // will be false if person is not in the map
		fmt.Println("Ann", "was at the meeting")
	}

	fmt.Println("Map content:")
	for key, value := range attendedSet {
		fmt.Printf("key: %s, value: %t\n", key, value)
	}
}
