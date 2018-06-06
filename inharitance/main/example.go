package main

import (
	"fmt"
)

type PanelCommons struct {
	// Modelname
	Name string
}

func (p *PanelCommons) printName() {
	fmt.Println(p.Name)
}

type RPPLine struct {
	// Common Panel Fields
	PanelCommons
	// add
	Variation string
}

func main() {
	rdb := RPPLine{
		PanelCommons: PanelCommons{
			Name: "6514",
		},
		Variation: "52-2",
	}
	rdb.printName()
	fmt.Println("")
	fmt.Println(rdb)
}
