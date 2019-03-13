package allocation

import "fmt"

func WithNew() {
	fmt.Println("=== Allocation with new ===")

	fmt.Printf("Empty int: %#+v\n", new(int))
	fmt.Printf("Empty Array: %#+v\n", new([]int)) // allocates slice structure; *p == nil; rarely useful
}
