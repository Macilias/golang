package main

import (
	"fmt"
	"strings"
)

func main() {

	arrays()
	slices()
	arrVSslices()
	arrVSslices2()
	zeroValue()
	useMakeOnSlices()
	twoDimensionalSlices()
	appendToSlice()
	growingSlices()
	filteringSlices()

}
func arrays() {
	fmt.Println("\narrays():")
	// Arrays
	var ar [2]string
	ar[0] = "Hello"
	ar[1] = "World"
	fmt.Println(ar[0], ar[1])
	fmt.Println(ar)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	fmt.Println("\nslices():")
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
}

func zeroValue() {
	fmt.Println("\nzeroValue():")
	/*
		The zero value of a slice is nil.
	*/
	var s2 []int
	fmt.Println(s2, len(s2), cap(s2))
	if s2 == nil {
		fmt.Println("nil!")
	}
	// whats about array? nope, can not convert a2 to nil:
	//var a2 [2]int
	//if a2 == nil {
	//	fmt.Println("nil!")
	//}
}

func arrVSslices() []int {
	fmt.Println("\narrVSslices():")
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
	printSlice(s)
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// Extend its length.
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	return s
}

func arrVSslices2() {
	fmt.Println("\narrVSslices2():")
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//           0    1    2    3    4    5
	s := []string{"g", "o", "l", "a", "n", "g"}
	printByteSlice(b)
	printByteSlice(b[1:4])
	printStringSlice(s)
	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
}

func useMakeOnSlices() {
	fmt.Println("\nuseMakeOnSlices():")

	a := make([]int, 5)
	printSlice(a, "a")

	b := make([]int, 0, 5)
	printSlice(b, "b")

	c := b[:2]
	printSlice(c, "c")

	d := c[2:5]
	printSlice(d, "d")

	e := make([]int, 5, 5)
	printSlice(e, "e")

	// len can not be larger than cap
	//f := make([]int, 5, 0)
	//printSliceWithName("f", f)
}

func twoDimensionalSlices() {
	fmt.Println("\ntwoDimensionalSlices():")
	// Create a tic-tac-toe board.
	board := [][]string{
		// I would leave the redundant declaration here, but it also indicates that
		// the second dimension is declared "inside", which is not the case!
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	fmt.Println("\nappendToSlice():")
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)

	printSlice(s)

	// and now append a whole array!!!
	a := make([]int, 5)
	a[0] = 5
	a[1] = 6
	a[2] = 7
	a[3] = 8
	a[4] = 9
	s = append(s, a...)

	printSlice(s)
}

func growingSlices() {
	fmt.Println("\ngrowingSlices():")
	// use append for automatic growing of slice.
	// here is an example how to do it manually:
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p == []byte{2, 3, 5, 7, 11, 13}
	printByteSlice(p)
}

func filteringSlices() {
	fmt.Println("\nfilteringSlices():")
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	printSlice(a, "after init")
	b := Filter(a, func(n int) bool { return n%2 == 0 })
	printSlice(b, "after filter")
}

// Helper Functions:

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func printSlice(s []int, n ...string) {
	if len(n) > 0 {
		fmt.Printf("%s len=%d cap=%d %v\n", n[0], len(s), cap(s), s)
		return
	}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printByteSlice(b []byte, n ...string) {
	if len(n) > 0 {
		fmt.Printf("%s len=%d cap=%d %v\n", n[0], len(b), cap(b), string(b[:]))
		return
	}
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b[:])
}

func printStringSlice(s []string, n ...string) {
	if len(n) > 0 {
		fmt.Printf("%s len=%d cap=%d %v\n", n[0], len(s), cap(s), s)
		return
	}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
