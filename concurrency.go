/*
Concurrent programming in many environments is made difficult by the subtleties
required to implement correct access to shared variables. Go encourages a different
approach in which shared values are passed around on channels and, in fact, never
actively shared by separate threads of execution. Only one goroutine has access to
the value at any given time. Data races cannot occur, by design. To encourage this
way of thinking we have reduced it to a slogan:

 - Do not communicate by sharing memory; instead, share memory by communicating. -

This approach can be taken too far. Reference counts may be best done by putting a
mutex around an integer variable, for instance. But as a high-level approach, using
channels to control access makes it easier to write clear, correct programs.

One way to think about this model is to consider a typical single-threaded program
running on one CPU. It has no need for synchronization primitives. Now run another
such instance; it too needs no synchronization. Now let those two communicate;
if the communication is the synchronizer, there's still no need for other synchronization.
Unix pipelines, for example, fit this model perfectly. Although Go's approach to concurrency
originates in Hoare's Communicating Sequential Processes (CSP), it can also be seen as a
type-safe generalization of Unix pipes.
*/

package main

import (
	"fmt"
	"runtime"
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

// === MULTI CHANNEL EXAMPLE ===

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

var MaxOutstanding = 5

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle(queue chan *Request) {
	for req := range queue { // TODO seems to unwrap the chan obj
		req.resultChan <- req.f(req.args)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // Wait to be told to exit. TODO when does this happen
}

/*
There's clearly a lot more to do to make it realistic, but this code is a framework
for a rate-limited, parallel, non-blocking RPC system, and there's not a mutex in sight.
*/
func MultiChannelExample() {
	fmt.Println("=== Multi Channel Example ===")
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	request2 := &Request{[]int{1, -9}, sum, make(chan int)}

	clientRequest := make(chan *Request)

	go Serve(clientRequest, make(chan bool))

	// Send request
	clientRequest <- request
	clientRequest <- request2

	// Wait for response.
	fmt.Printf("1. answer: %d\n", <-request.resultChan)
	fmt.Printf("2. answer: %d\n", <-request2.resultChan)
}

// === PARALLELIZATION ===
/*
Another application of these ideas is to parallelize a calculation across multiple CPU cores.
If the calculation can be broken into separate pieces that can execute independently, it can
be parallelized, with a channel to signal when each piece completes.
*/

type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

var _ = runtime.NumCPU() // number of CPU cores
// reports (or sets) the user-specified number of cores that a Go program can have running simultaneously.
var numCPU = runtime.GOMAXPROCS(0) // just queries with 0 value

// We launch the pieces independently in a loop, one per CPU.
// They can complete in any order but it doesn't matter;
// we just count the completion signals by draining the channel
// after launching all the goroutines.
func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}
	// All done.
}

func (v Vector) Op(f float64) float64 {
	return f
}

// === MAIN ===
func main() {
	GoroutineExample()
	SimpleChannelExample()
	MultiChannelExample()
	var vec1, vec2 Vector
	vec1 = []float64{1, 2, 3}
	vec2 = []float64{3, 2, 1}
	vec1.DoAll(vec2)
	fmt.Printf("Result: %v", vec1)
}
