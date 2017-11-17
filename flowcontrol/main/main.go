package main

import (
	"fmt"
	"runtime"
)

func main() {
	forStyle()
	whileStyle()
	ifFlavours()
	swichTest()
}

// Loops

func forStyle() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func whileStyle() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// If

func ifFlavours() {
	x, y := 2, 1

	// simple
	if x < y {
		// the word end is near
	}

	// with short statement
	if z := x + y; z == y {
		// x must be 0
	} else {
		// x ist something else
		// Variables declared inside an if short statement are also available
		// inside any of the else blocks.
	}
}

func swichTest() {
	// all like in Java but:
	// - Go only runs the selected case, not all the cases that follow.
	//   In effect, the break statement that is needed -> equivalent with *fallthrough*
	//   The 'fallthrough' must be the last thing in the case!
	// - Go's switch cases need not be constants, and the values involved need not be integers.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
