package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	t1 := 0
	t2 := 1
	return func() int {
		res := t1
		t1 = t2
		t2 = res + t2
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
