package simplemath

import(
	"errors"
)

// To make public and private a function you just need to Capitalize the function name

// if you have parameters of the same type, you declare them both then the type for both
func Add(p1, p2 float64) float64 {
	return p1 + p2
}

// you can name the returning values, but its preference beacuse some people think its harder to read
func Divide(p1 , p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}
	answer = p1 / p2
	return
}

// variadic parameters function
func Sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {	
		total += value 
	}
	return total
}

