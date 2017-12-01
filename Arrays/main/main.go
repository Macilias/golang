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
	var sl []int = primes2[1:4]
	fmt.Println(sl)

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

	/*
		This is an array literal:
		[3]bool{true, true, false}

		And this creates the same array as above, then builds a slice that references it:
		[]bool{true, true, false}

		For the array: var a [10]int
		these slice expressions are equivalent:

		a[0:10]
		a[:10]
		a[0:]
		a[:]

		A slice has both a length and a capacity.
		The length of a slice is the number of elements it contains.
		The capacity of a slice is the number of elements in the underlying array,
		counting from the first element in the slice.
	*/

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("s", s)

	// Extend its length.
	s = s[:4]
	printSlice("s", s)

	// Drop its first two values.
	s = s[2:]
	printSlice("s", s)

	/*
		The zero value of a slice is nil.
	*/

	var s2 []int
	fmt.Println(s2, len(s2), cap(s2))
	if s == nil {
		fmt.Println("nil!")
	}

	// whats about array? nope, can not convert a2 to nil:
	//var a2 [2]int
	//if a2 == nil {
	//	fmt.Println("nil!")
	//}

	// using make on slices
	useMakeOnSlices()
}

func useMakeOnSlices() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	e := make([]int, 5, 5)
	printSlice("e", e)

	// len can not be larger than cap
	//f := make([]int, 5, 0)
	//printSlice("f", f)
}

func printSlice(n string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", n, len(s), cap(s), s)
}
