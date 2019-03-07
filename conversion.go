package main

import (
	"fmt"
	"sort"
)

// ### Conversions ###

type Sequence []int

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort() // It's an idiom to convert the type to access a different set of methods
	return fmt.Sprint([]int(s))
}

// ### Interface conversions and type assertions ###

type Stringer interface {
	String() string
}

type StringerType string

type CopyableStringer interface {
	String() string
	Copy() Sequence
}

func isConvertibleToString(value interface{}) string {
	switch str := value.(type) {
	case string:
		return "string: " + str
	case Stringer:
		return "Stringer: " + str.String()
	}
	return "not valid"
}

func main() {
	sequences := Sequence{1, 7, 3, 2, 9}
	fmt.Println(sequences)
	fmt.Println(isConvertibleToString("ImAString"))
	fmt.Println(isConvertibleToString(sequences))
	fmt.Println(isConvertibleToString(2))

	var stringer CopyableStringer = sequences

	str, ok := stringer.(Stringer) // type assertions work for interfaces only
	if ok {
		fmt.Printf("convertible: %q\n", str)
	} else {
		fmt.Printf("not convertible\n")
	}

	//var stringerType StringerType = "ddfdf"
	//str = string(str) // type assertions work for interfaces only
	//fmt.Printf("convertible: %q\n", str)

	// Some interface checks do happen at run-time, though.
	// One instance is in the encoding/json package, which defines a Marshaler interface.
	// When the JSON encoder receives a value that implements that interface,
	// the encoder invokes the value's marshaling method to convert it to JSON instead of doing the
	// standard conversion. The encoder checks this property at run time with a type assertion like:
	// var val = ""
	// m, ok := val.(json.Marshaler) // TODO is not at run-time for me
}
