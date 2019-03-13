package main

import "github.com/golang-tests/allocation"

func main() {
	allocation.WithNew()
	allocation.WithMake()
}
