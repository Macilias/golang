package main

import (
	"fmt"
	"math"
)

func main() {
	functionValues()
	functionClosures()
}

func functionValues() {
	fmt.Println("\nfunctionValues():")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	// the result of direct usage of the named function is the same as,...
	fmt.Println(hypot(3, 4))
	// ... its result, when passed to our compute function, which enters same values.
	fmt.Println(compute(hypot))
	// the compute with 3,4 function can also call other methods with same signature.
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/*
	Go functions may be closures. A closure is a function value that references variables from
	outside its body. The function may access and assign to the referenced variables; in this sense
	the function is "bound" to the variables.
*/
func functionClosures() {
	fmt.Println("\nfunctionClosures():")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
