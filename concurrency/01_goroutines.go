package concurrency

import (
	"fmt"
	"time"
)

// ### GOROUTINE ###
/*
A goroutine has a simple model:
it is a function executing concurrently with other goroutines in the same address space.
It is lightweight, costing little more than the allocation of stack space. And the stacks
start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while
waiting for I/O, others continue to run. Their design hides many of the complexities of
thread creation and management.
*/

/*
Prefix a function or method call with the go keyword to run the call in a new goroutine.

When the call completes, the goroutine exits, silently. (The effect is similar to the
Unix shell's & notation for running a command in the background.)

	go list.Sort()  // run list.Sort concurrently; don't wait for it.

*/

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

func GoroutineExample() {
	Announce("my message from the goroutine", 0) // TODO why does this occur in the stdout without channel?
	// In Go, function literals are closures: the implementation makes sure the variables
	// referred to by the function survive as long as they are active.
	//
	//These examples aren't too practical because the functions have no way of
	// signaling completion. For that, we need CHANNELS.
}
