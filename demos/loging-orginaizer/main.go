package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	// there are no exeptions in go , we just return error values
	f, err := os.Open("myapp.log")
	// nil is simmilar that null but its not a null value like other lenguajes
	if err != nil {
		log.Fatal(err) 
	}
	// we canot declare somthing and dont use it
	// deferes refers that after the main funcion is closed its goint to go and  close these ones at the end
	defer f.Close()
	
	// reading the input file from a network or somthing else
	r := bufio.NewReader(f)
	// there are no loops others than for loops, for loops are for all the cases and there are just 4 scenarios
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "ERROR") {
			fmt.Println(s) 
		}
	}
}