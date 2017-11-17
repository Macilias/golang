package main

import (
	"fmt"
	"math"
)

// Variables with initializers
var i, j int = 1, 2

// PS Outside a function, every statement begins with a keyword (var, func, and so on)
// and so the := construct is not available.

func main() {
	// Inside a function, the := short assignment statement can be used in place of a var
	// declaration with implicit type.
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
	defaultValues()
	typeConversion()
}

func defaultValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// The expression T(v) converts the value v to the type T.
func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

/*
>>>>> Basic types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide
on 64-bit systems. When you need an integer value you should use int unless you have a specific
reason to use a sized or unsigned integer type.

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

>>>> Type inference

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

>>>>> Constants

Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.

*/
