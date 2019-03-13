package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"syscall"
)

// === ERROR ===

// By convention, errors have type error, a simple built-in interface:
//
// type error interface {
//	Error() string
// }

// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// PathError's Error generates a string like this:
// 		open /etc/passwx: no such file or directory

func ErrorExample() {
	for try := 0; try < 2; try++ {
		_, err := os.Create("testfile")
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			// deleteTempFiles()  // Recover some space.
			continue
		}
		return
	}
}

// === PANIC ===

// What if he error is unrecoverable? Sometimes the program simply cannot continue.
// For this purpose, there is a built-in function panic that in effect creates a
// run-time error that will stop the program

// A toy implementation of cube root using Newton's method.
const USER = "ME"

func init() {
	if USER == "" {
		panic("no value for $USER")
	}
}

// This is only an example but real library functions should avoid panic.
// If the problem can be masked or worked around, it's always better to let
// things continue to run rather than taking down the whole program.
// One possible counterexample is during initialization: if the library
// truly cannot set itself up, it might be reasonable to panic, so to speak.

// === RECOVER ===

// Error is the type of a parse error; it satisfies the error interface.
type Error string

func (e Error) Error() string {
	return string(e)
}

type Work string

// === EXAMPLE: SAVE GOROUTINE ERROR ===
/*
One application of recover is to shut down a failing goroutine inside
a server without killing the other executing goroutines.
*/
func server(workChan <-chan *Work) {
	for work := range workChan {
		wg.Add(1)
		go safelyDo(work)
	}
	fmt.Println("waiting for goroutines to finish")
	wg.Wait()
}

func safelyDo(work *Work) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed but recovered from it:", err)
		}
	}()
	fmt.Println(*work)
	panic(Error("PANIC!!!"))
}

type Regexp regexp.Regexp

func (regexp *Regexp) doParse(str string) *Regexp {
	return regexp
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
	panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
	regexp = new(Regexp)
	// doParse will panic if there is a parse error.
	defer func() {
		if e := recover(); e != nil {
			regexp = nil    // Clear return value.
			err = e.(Error) // Will re-panic if not a parse error.
		}
	}()
	return regexp.doParse(str), nil
}

var wg sync.WaitGroup // main would kill all goroutines without that

func main() {
	var workChannel = make(chan *Work, 2)
	one := Work("one")
	two := Work("two")
	workChannel <- &one
	workChannel <- &two
	close(workChannel) // loop would never stop reading from the channel
	server(workChannel)
	fmt.Println("main finished!")
}
