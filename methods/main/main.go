package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// a method is just a function with a receiver argument.
func main() {
	receiverMethod()
	methodAsFunction()
	methodsOnNumericType()
	pointerReceivers()
	pointerFunctions()
}

func receiverMethod() {
	fmt.Println("\nreceiverMethod():")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodAsFunction() {
	fmt.Println("\nmethodAsFunction():")
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// this is a little bit of a cheat, but it should point out that also simple types can be used
// as receiver and also not ONCE THAT HAS BEEN DECLARED WITHIN THIS FILE (but must be in package!)
func methodsOnNumericType() {
	fmt.Println("\nmethodsOnNumericType():")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
	With a value receiver, the Scale method operates on a copy of the original Vertex value.
	(This is the same behavior as for any other function argument.)
	The Scale method must have  a pointer receiver to change the Vertex value declared in the main function.
*/
func pointerReceivers() {
	fmt.Println("\npointerReceivers():")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}

// WHEN YOU REMOVE THE POINTER * HERE, v want be scaled, but its copy will be scaled instead.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerFunctions() {
	fmt.Println("\npointerFunctions():")
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/*
	This case (func with argument) is a bit different from the case before (method with receiver).
	If you use a pointer here you must pass it with &v otherwise you get a compile error.
*/
// HERE YOU CAN NOT REMOVE THE *
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
