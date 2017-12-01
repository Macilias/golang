package main

import "fmt"

// x and y are named here so they are treated as variables defined at the top of the function.
// A return statement without arguments returns the named return values.
// This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
