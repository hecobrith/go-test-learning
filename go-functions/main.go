package main

import (
"fmt"
"github.com/hecobrith/go-test-learning/go-functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(6, 0)
	if err != nil {
		fmt.Printf("An error ocurred %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}
	fmt.Printf("%f\n" , simplemath.Add(2.244, 6.1214))
	numbers:= []float64{12.13, 123.13, 123.12, 315.43}
	total := simplemath.Sum(numbers...)
	fmt.Printf("total of sum: %f\n", total)

	sv := simplemath.NewSemanticVersion(1,2,3)
	fmt.Println(sv.String())

	sv.IncrementMayor()
	println(sv.String())
	sv.IncrementMayor()
	println(sv.String())

	// anonimous functions
	func() {
		println("stuff")
	}()

	// decalred function like an arrow functions
	a := func(name string) string {
		fmt.Printf("some anonimous %s function \n", name)	
		return name	
	}

	a("functuon 1")
	a("functuon 2")
	value := a("function 3")
	println(value)
}
