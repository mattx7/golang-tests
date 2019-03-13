/*
Creates slices, maps, and channels only,
and it returns an initialized (not zeroed) value of type T (not *T)
*/

package allocation

import "fmt"

func WithMake() {
	fmt.Println("=== Allocation with make ===")

	fmt.Printf("Slice: %#+v\n", make([]int, 3, 10)) // the slice v now refers to a new array of 10 ints
}
