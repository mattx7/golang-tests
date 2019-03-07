/*
Creates slices, maps, and channels only,
and it returns an initialized (not zeroed) value of type T (not *T)
*/
package main

import "fmt"

func main() {
	fmt.Printf("Empty int: %#+v\n", new(int))
	fmt.Printf("Empty Array: %#+v\n", new([]int)) // allocates slice structure; *p == nil; rarely useful
	fmt.Printf("Slice: %#+v\n", make([]int, 3, 10)) // the slice v now refers to a new array of 10 ints

}
