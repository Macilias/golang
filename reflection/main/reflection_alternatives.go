package main

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/kr/pretty"
)

type foo struct {
	Desc1, Desc2, Desc3 string
}

type bar struct {
	Desc    string
	Model   string
	Page    string
	STPreis []RDBSTPreis
}

type RDBSTPreis struct {
	Desc string `json:"seitenteilpreis_beschreibung"`
	Pvc  string `json:"pvc"`
	Alu  string `json:"alu"`
}

type bars struct {
	list []bar
}

var fooFixt = `[
	{
		"Page": "1",
		"Desc1":"beschreibung1",
		"Desc2":"beschreibung2",
		"Desc3":"beschreibung3",
		"Model1": "modell1",
		"Model2": "modell2",
		"Model3": "modell3",
        "seitenteilpreis1": null,
        "seitenteilpreis2": [
            {
                "seitenteilpreis_beschreibung": "Seitenteilpreis in % (von der jew. Ausf\u00fchrung der Haust\u00fcrf\u00fcllung inkl. Glas)",
                "pvc": "75%",
                "alu": "75%"
            }
        ],
        "seitenteilpreis3": null
	}
]`

func MustStr(m json.RawMessage) string {
	var out string
	err := json.Unmarshal(m, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func (bars *bars) UnmarshalJSON(b []byte) error {
	data := []map[string]json.RawMessage{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	for _, d := range data {
		for i := 1; i <= 3; i++ {
			var stpreis []RDBSTPreis
			err := json.Unmarshal(d[fmt.Sprintf("seitenteilpreis%d", i)], &stpreis)
			if err != nil {
				panic(err)
			}
			bars.list = append(bars.list, bar{
				Desc:    MustStr(d[fmt.Sprintf("Desc%d", i)]),
				Model:   MustStr(d[fmt.Sprintf("Model%d", i)]),
				Page:    MustStr(d["Page"]),
				STPreis: stpreis,
			})
		}
	}
	return nil
}

func main() {
	var bars bars
	err := json.Unmarshal([]byte(fooFixt), &bars)
	if err != nil {
		logrus.Fatal(err)
	}

	pretty.Print(bars.list)
	fmt.Printf("\n\n")
}

func alt() {
	var foos []map[string]interface{}
	err := json.Unmarshal([]byte(fooFixt), &foos)
	if err != nil {
		logrus.Fatal(err)
	}

	pretty.Print(foos[0])
	fmt.Printf("\n\n")
	aus := foos[0]["ausfuehrung2"].([]interface{})
	einser := aus[0].(map[string]interface{})
	pretty.Print(einser["pvc"])
	fmt.Printf("\n\n")
}

/*
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
