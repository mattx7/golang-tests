/*
Go does not provide the typical, type-driven notion of subclassing,
but it does have the ability to “borrow” pieces of an implementation
by embedding types within a struct or interface.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
// It is a union of the embedded interfaces (which must be disjoint sets of methods).
type ReadWriter interface {
	Reader
	Writer
}

// ### The same basic idea applies to structs, but with more far-reaching implications.

// ReadWriter stores pointers to a Reader and a Writer.
// The embedded elements are pointers to structs and of course
// must be initialized to point to valid structs before they can be used.
// The ReadWriter struct could be written as
type BufioReadWriter struct {
	*bufio.Reader // dont provide field names to embed struct and avoid redefining the methods todo check that
	*bufio.Writer
}

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Printf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

func main() {
	job := &Job{"RUN_THINGS", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Println("starting now...")
	job.Printf("%#v", job)
}

/*
Embedding types introduces the problem of name conflicts but the rules to resolve them are simple.
First, a field or method X hides any other item X in a more deeply nested part of the type.
If log.Logger contained a field or method called Command, the Command field of Job would dominate it.

Second, if the same name appears at the same nesting level, it is usually an error;
it would be erroneous to embed log.Logger if the Job struct contained another field
or method called Logger. However, if the duplicate name is never mentioned in the
program outside the type definition, it is OK. This qualification provides some protection
against changes made to types embedded from outside; there is no problem if a field
is added that conflicts with another field in another subtype if neither field is ever used.
*/
