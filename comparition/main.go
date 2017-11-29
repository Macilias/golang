package main

import (
	"fmt"
)

func main() {

	// This is probably a bug!
	po := 3.141592653589793
	pi := 3.14159265358979323846264338327950288419716939937510582097494459
	if po == pi {
		fmt.Println("this should not be true, but is")
	}
}
