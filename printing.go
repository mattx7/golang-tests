package main

import (
	"fmt"
	"os"
)
type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

type MyString string

//func (m MyString) String() string {
//	return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
//}

func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}

func main() {

	fmt.Printf("\n==> Different Prints:\n")
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Printf("\n==> Print dezimals and hexadezimals:\n")
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))


	fmt.Printf("\n==> Print map:\n")
	timeZone := map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Printf("%v\n", timeZone) // or just fmt.Println(timeZone)


	fmt.Printf("\n==> Print structs:\n")

	t := &T{ 7, -2.35, "abctdef" }
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Printf("\n==> Default format for a custom type:\n")
	fmt.Printf("%v\n", t)
}
