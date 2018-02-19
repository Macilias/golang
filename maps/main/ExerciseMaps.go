package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)

	f := strings.Fields(s)
	for _, value := range f {
		m[value]++
	}
	return m
}

// func main() {
// 	wc.Test(WordCount)
// }
