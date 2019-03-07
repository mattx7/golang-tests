package main

import "fmt"

const Enone = 0
const Eio = 1
const Einval = 2

type File struct {
	fd      int
	name    string
	dirinfo string
	nepipe  int
}

func NewFileV1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = ""
	f.nepipe = 0
	return f
}

func NewFileV2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name, "", 0} // composite literal
}

func NewFileV3(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{name: name, fd: fd} // labeled elements
}

func main() {
	fmt.Println("===== Composed Literals ====")

	fmt.Printf("%+v\n", NewFileV1(0, "normal assignment"))
	fmt.Printf("%#v\n", NewFileV2(0, "composite literals"))
	fmt.Println(NewFileV3(0, "with labeled parameter"))

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Printf("Array: %#v\n", a)
	fmt.Printf("Splice: %#v\n", s)
	fmt.Printf("Map: %#v\n", m)

	fmt.Println("============================")
}
