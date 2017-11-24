package main

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
