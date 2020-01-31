package main

import (
	"fmt"
)

func main() {
	mySlice:= make([]int, 1, 4)
	fmt.Printf("Lenght is: %d capacity is : %d", 
		len(mySlice), cap(mySlice))

	for i:= 1; i < 17 ; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Capacity is : %d", 
		cap(mySlice))
	}
}