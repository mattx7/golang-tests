package concurrency

import "fmt"

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
