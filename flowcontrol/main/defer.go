package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Defer (pushes a function call onto a list. - use like finally in Java)

func copyTest() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	CopyFile("dat2", "dat1")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	check(err)
	defer src.Close()

	dst, err := os.Create(dstName)
	check(err)
	defer dst.Close()

	return io.Copy(dst, src)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
	Defer statements allow us to think about closing each file right after opening it,
	guaranteeing that, regardless of the number of return statements in the function,
	the files will be closed.
	There are 3 simple rules:
	1. A deferred function's arguments are evaluated when the defer statement is evaluated.
	2. Deferred function calls are executed in Last In First Out order after the surrounding
		function returns.
	3. Deferred functions may read and assign to the returning function's named return values.
*/

func deferTest() {
	fmt.Println(">>>>>> rule1: <<<<<<")
	rule1()
	fmt.Println(">>>>>> rule2: <<<<<<")
	rule2()
	fmt.Println("\n>>>>>> rule3: <<<<<<f")
	fmt.Println(rule3())
}

// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func rule1() {
	i := 0
	// this is called after the function returns
	defer fmt.Println("defered print", i)
	i++
	fmt.Println("normal print", i)
	return
}

// 2. Deferred function calls are executed in LIFO order after the surrounding function returns.
func rule2() {
	// LIFO: i=0,1,2,3 -> 3210
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 3. Deferred functions may read and assign to the returning function's named return values.
func rule3() (i int) {
	defer func() { i++ }()
	return 1
}

// Other uses of defer (beyond the file.Close example given earlier) include
/*
releasing a mutex:
mu.Lock()
defer mu.Unlock()

printing a footer:
printHeader()
defer printFooter()
and more.
*/
