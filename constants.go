/*
Constants in Go are just that—constant. T
hey are created at compile time, even when defined as locals in functions, and can only be numbers, characters (runes),
strings or booleans. Because of the compile-time restriction, the expressions that define them must be constant
expressions, evaluatable by the compiler. For instance, 1<<3 is a constant expression, while math.Sin(math.Pi/4) is
not because the function call to math.Sin needs to happen at run time.
*/
package main

import (
	"fmt"
	"os"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	MON uint8 = iota + 1 // type uint8 and byte is the same
	TUE
	// WEN
	THU = iota + 2
	FRI
	SAT
	SUN
)

// does not work!
//const PI4 = math.Sin(math.Pi/4)

// Variables can be initialized just like constants but the initializer can be a general expression computed at run time.
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
	fmt.Println(MON, TUE, THU, FRI, SAT, SUN)
	fmt.Println(home, user, gopath)

}
