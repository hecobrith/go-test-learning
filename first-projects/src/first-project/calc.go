package main

import (
	"fmt"
	"reflect"
)

func main() {
	// short asignment form only work inside functions scope and its the idiomatic way
	a := 10.0000
	b := 3

	fmt.Println("\na is type of", reflect.TypeOf(a), "and b is type of", reflect.TypeOf(b))
	// here we didnt changed the variable type, it still continue to be a float, its just used like that on the calculation
	c := int(a) + b

	fmt.Println( "\nC has a value:", c, "and is type", reflect.TypeOf(c))
}