/*
There are major differences between the ways ARRAYS work in Go and C. In Go,
- Arrays are values. Assigning one array to another COPIES all the elements.
- In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
- The size of an array is part of its type. The types [10]int and [20]int are DISTINCT.
*/
package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

// byte = uint8
func customAppend(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

// Go's arrays and slices are one-dimensional.
// To create the equivalent of a 2D array or slice,
// it is necessary to define an array-of-arrays or slice-of-slices, like this:
type Transform [3][3]float64 // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte    // A slice of byte slices.

func main() {

	fmt.Println("\n==> basic arrays")
	array1 := [3]float64{7.0, 8.5, 9.1}
	array2 := [...]float64{7.0, 8.5, 9.1}
	fmt.Printf("%#+v\n", Sum(&array1))                                         // Note the explicit address-of operator
	fmt.Printf("%#+v\n", Sum(&array2))                                         // Note the explicit address-of operator
	fmt.Println("But even this style isn't idiomatic Go. Use slices instead!") // Note the explicit address-of operator

	fmt.Println("\n==> basic slices")
	basicSlice := []byte{0, 1} // this is a slice and not an array
	// func make([]T, len, cap) []T
	basicSlice = make([]byte, 2, 2)
	basicSlice[0] = 0
	basicSlice[1] = 1
	bytes := customAppend(basicSlice, []byte{254, 255}) // 255 max because byte/uint8
	fmt.Println(bytes)

	fmt.Println("\n==> build-in append")
	x := []int{1, 2, 3}
	y := append(x, 4, 5, 6)
	fmt.Println(y)
	fmt.Printf("x:%p \ny:%p\n", &x, &y)

	x = []int{1, 2, 3}
	y = []int{4, 5, 6}
	x = append(x, y...) // "..." unwraps a slice or something like that
	fmt.Println(x)

	fmt.Println("\n==> Two-dimensional slices")
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	//fmt.Println(text)
	for idx, line := range text {
		fmt.Printf("[%d] %s\n", idx, line)
	}

}
