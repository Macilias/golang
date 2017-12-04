package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	directAccess()
	mapLiterals()
	mutatingMaps()
}

func directAccess() {
	fmt.Println("\ndirectAccess():")
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Nimmerland"])
}

func mapLiterals() {
	fmt.Println("\nmapLiterals():")
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Nimmerland"])
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v1, ok := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v2, ok := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok)
}
