package main

import (
	"fmt"
	"runtime"
	"time"
)

// Switch

func switchTest() {
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

func switchWithNoConditionTest() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
