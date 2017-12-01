package main

import "fmt"

func main() {

	// Arrays
	var ar [2]string
	ar[0] = "Hello"
	ar[1] = "World"
	fmt.Println(ar[0], ar[1])
	fmt.Println(ar)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	primes2 := [6]int{2, 3, 5, 7, 11, 13}

	// Slices are like references to arrays
	/*
		A slice does not store any data, it just describes a section of an underlying array.
		Changing the elements of a slice modifies the corresponding elements of its underlying array.
		Other slices that share the same underlying array will see those changes.
	*/
	var s []int = primes2[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
