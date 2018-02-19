package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
)

type foo struct {
	Desc1, Desc2, Desc3 string
}

var fooFixt = `[
	{ 
		"Desc1":"beschreibung1",
		"Desc2":"beschreibung2", 
		"Desc3":"beschreibung3"
	}
]`

func main() {
	var foos []foo
	err := json.Unmarshal([]byte(fooFixt), &foos)
	if err != nil {
		logrus.Fatal(err)
	}

	parseFoo4(descGenerator(1, foos[0]))
	parseFoo4(fieldGenerator(2, "Desc", foos[0]))
	fmt.Printf("\n\n")
}

func descGenerator(i int, f foo) func() string {
	return func() string {
		v := reflect.ValueOf(f)
		return v.FieldByName(fmt.Sprintf("Desc%d", i)).String()
	}
}

func descBlaGenerator(i int, f foo) (func() string, func() string) {
	return fieldGenerator(i, "Desc", f), fieldGenerator(i, "Bla", f)
}

func descBlaGenerator2(i int, f foo) (func() string, func() string) {
	return func() string {
			v := reflect.ValueOf(f)
			return v.FieldByName(fmt.Sprintf("Desc%d", i)).String()
		}, func() string {
			v := reflect.ValueOf(f)
			return v.FieldByName(fmt.Sprintf("Bla%d", i)).String()
		}
}

func fieldGenerator(i int, field string, f foo) func() string {
	return func() string {
		v := reflect.ValueOf(f)
		return v.FieldByName(fmt.Sprintf("%s%d", field, i)).String()
	}
}

func parseFoo4(getDesc func() string) {
	fmt.Printf("%#v", getDesc())
}

/*
func parseFoo(f foo) {
	fmt.Printf("%v %v %v\\nn", f.Desc1, f.Desc2, f.Desc3)
}

func parseFoo2(f foo) {
	v := reflect.ValueOf(f)
	for i := 1; i <= 3; i++ {
		fmt.Printf("%v\n", v.FieldByName(fmt.Sprintf("Desc%d", i)).String())
	}
}
*/
