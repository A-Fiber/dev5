package main

import (
	"fmt"
)

//defer

/*
func main() {
	var pI *int //memory address ==> of a value of type int
	var I int = 3
	pI = &I
	increment(pI)
	fmt.Println(I)
	fmt.Println(pI)
}

func increment(pI *int) {
	*pI++ //de-referencing
}
*/

//defer is LIFO
/*
func main() {
	defer fmt.Println("Hello1")
	defer fmt.Println("Hello2")
	fmt.Println("World")
}
*/

//panics the equivalent of throwing an exception in other languages
/*
func main() {
	fmt.Println("Hello")
	testpanics()
	fmt.Println("World")
}

func testpanics() {
	panic("A panic happened")
}
*/

//to recover from a panic use the defer statement with a powerful feature named
//anonymous functions

func main() {
	fmt.Println("Hello")
	testpanics()
	fmt.Println("World")
}

func testpanics() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("We recovered from a panic")
		}
	}()
	panic("A panic happened")
}
