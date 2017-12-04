package main

import "fmt"

func main() {

	simpleExample()

}

func simpleExample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}
