package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sayHello(s string) string {
	return "Hello " + s
}

func main() {
	var s string = "true"
	var b bool = true
	var f float64 = 11.222

	var beatles [4]string

	beatles[0] = "John"
	beatles[1] = "Paul"
	beatles[2] = "Ringo"
	beatles[3] = "George"

	// checking types
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(f))

	// type convertion
	fmt.Println(strconv.ParseBool(s))
	fmt.Println(reflect.TypeOf(strconv.FormatBool(true)))
	fmt.Println(strconv.ParseInt("22", 10, 32))
}
