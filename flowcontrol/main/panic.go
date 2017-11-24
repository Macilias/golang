package main

import "fmt"

// Panic
// seams to be used similar to try catch in Java, consider following example:
/*
	In th json package from the Go standard library.
	It decodes JSON-encoded data with a set of recursive functions.
	When malformed JSON is encountered, the parser calls panic to unwind the stack
	to the top-level function call,
	which recovers from the panic and returns an appropriate error value
	(see the 'error' and 'unmarshal' methods of the decodeState type in decode.go).
*/

func panicTest() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	//This is unnecessary, but a reminder that defers are between return and recover
	//defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Prognose:

/*
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4 // this was the i during panic, must have been passed over to recover
Returned normally from g.
*/
