package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// if you need to imitate onSet behaviour from different languages
func (v *Vertex) SetX(x int) {
	// do your magic checks here
	v.X = x
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 3
	fmt.Println(v)
	v.SetX(4)
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
