package main

import (
	"fmt"
	"time"
)

func main() {
	goRoutines()
	channels()
	bufferedChannels()
	rangeAndClose()
}

/*

>>>> Goroutines
	go f() starts f in a new goroutine
	A goroutine is a lightweight thread managed by the Go runtime.

	go f(x, y, z)
	starts a new goroutine running

	f(x, y, z)
	The evaluation of f, x, y, and z happens in the current goroutine and the
	execution of f happens in the new goroutine.

	Goroutines run in the same address space, so access to shared memory must be synchronized.
	The sync package provides useful primitives.

 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goRoutines() {
	fmt.Println("\ngoRoutines:")
	go say("world")
	say("hello")
}

/*

>>>> Channels
	Channels are a typed conduit (Leitung) through which you can send and receive values with the
	channel operator, <-.

	ch <- v    // Send v to channel ch.
	v := <-ch  // Receive from ch, and
			   // assign value to v.
	(The data flows in the direction of the arrow.)

	Like maps and slices, channels must be created before use:

	ch := make(chan int)
	By default, sends and receives block until the other side is ready.
	This allows goRoutines to synchronize without explicit locks or condition variables.

*/

// The following example code sums the numbers in a slice,
// distributing the work between two goRoutines.
// Once both goRoutines have completed their computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channels() {
	fmt.Println("\nchannels:")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	// TODO understand following:
	// why does c puts its inside into y first and into x the second time?

	fmt.Println(x, y, x+y)
}

/*

>>>> Buffered Channels
	Channels can be buffered. Provide the buffer length as the second argument to make to initialize
	a buffered channel:

	ch := make(chan int, 100)

	Sends to a buffered channel block only when the buffer is full.
	Receives block when the buffer is empty.

 */

func bufferedChannels() {
	fmt.Println("\nbufferedChannels:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*

>>>> Range and Close
	A sender can close a channel to indicate that no more values will be sent.
	Receivers can test whether a channel has been closed by assigning a second parameter to the
	receive expression: after

	v, ok := <-ch
	ok is false if there are no more values to receive and the channel is closed.

	The loop for i := range c receives values from the channel repeatedly until it is closed.

	Note: Only the sender should close a channel,
	never the receiver. Sending on a closed channel will cause a panic.

	Another note: Channels aren't like files; you don't usually need to close them.
	Closing is only necessary when the receiver must be told there are no more values coming,
	such as to terminate a range loop.

 */

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndClose() {
	fmt.Println("\nrangeAndClose:")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	i := 1
	for f := range c {
		fmt.Println(i, f)
		i++
	}
}
