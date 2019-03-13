package concurrency

import "fmt"

// ### CHANNELS ###

/*
Like maps, channels are allocated with make, and the resulting value acts as a reference
to an underlying data structure. If an optional integer parameter is provided,
it sets the buffer size for the channel. The default is zero, for an unbuffered or synchronous channel.

ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files

Unbuffered channels combine communication—the exchange of a value—with synchronization—guaranteeing
that two calculations (goroutines) are in a known state.
*/

// === SIMPLE CHANNEL EXAMPLE ===

func sumWithChan(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func SimpleChannelExample() {
	fmt.Println("=== Simple Channel Example ===")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sumWithChan(s[:len(s)/2], c)
	go sumWithChan(s[len(s)/2:], c)

	x, y := <-c, <-c
	// receive from c
	fmt.Println(x, y, x+y)
}
