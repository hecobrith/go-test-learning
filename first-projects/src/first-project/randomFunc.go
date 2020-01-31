package main

import (
	"math/rand"
	"time"
	"fmt"
)

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}

func main(){
	randomTime := random()
	fmt.Println(randomTime)
}