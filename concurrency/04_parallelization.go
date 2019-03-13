package concurrency

import (
	"fmt"
	"runtime"
)

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

func ParallelizationExample() {
	fmt.Println("=== ParallelizationExample ===")
	var vec1, vec2 Vector
	vec1 = []float64{1, 2, 3}
	vec2 = []float64{3, 2, 1}
	vec1.DoAll(vec2)
	fmt.Printf("Result: %v", vec1)
}
