# GO

## Enviormental variabls

The os Library have most of the stuff you need for any enviormental stuff

This is a simple function to print all the environ variables.
Olso with the function ```os.Getenv("USERNAME")``` you can get specific env variables

```go
import {
    "os"
    "fmt"
}

func main(){
    for _, env:= range os.Environ() {
        fmt.Println(env)
    }
}
```

## Functions

As other lenguajes they are first class citizens ergo they are functional lenguaje oriented. Syntax is easy and usual:

```go
func <name>(<params><type>)<return type> {
    <block of code>
}
```

Parameters always need what type of value they recived and after that what type of value the function return as the following example

```go
func titleCase(text string) string {
    fmt.Println(text)
    return text
}
```

## Conditions & Branching

The conditional always need to be boolean, can valuate strings or ints, and works witha usual operators. It suply nesting.
Also swiches is aother way of branching the syntax is olmos the same

```go
if <boolean exp> {
    <code block>
} else if {
    <code block>
} else {

}
```

Error handling is the most commun use case for conditions

## Loops

Go only uses for loops of three types: infinite (while), range, conditional.

```go
for <expresion> {
    <code>
}

for {
    <code>
}

for i := range <list or number> {
    <code>
}
```

## Arrays & Slices

"Named lists containing elements of the same type". Arrays fot fixed lenght when created (cos memory alocation), so for finamic ranges you need to use Slices, The thing is Slices in essence are arrays managed by Go, when you declare a Slice in reality Go creates an array. These meaning, that slices are always representation of memory alogation, so if you change the data on a slice will **ALWAYS** be reflected on the array. What Go does is creating a Slice inside an array, that when it reach the limit it doubles the size of the array, for more reference run the funcion of array expansión. creating arrays need the ```make()``` function in order to make new slices.

```go
make(<type>, <length>, <capacity>)
len(<slice>) // will always return the size of slice
cap() // will return the capacity of any slice

// there are other ways for creating slices
slice := [] int {1,2,3,4,5,6}
```

So if you select a pice of an array it becomes an slice, meaning you can select similar to other lenguajes ```slice[2:5] ```beeing the first number the inclusive index, and the second number the exolusive index. In order to make an array bigger we need to convert it into a slice first so... ``` slice.append()``` it's posible while ``` array.append()``` is not valid. But if a slice is an array in escence how does it deal with adding more space, if the slice have dinamyc capacity?? thats by making memory alocations of the len(array)^2, Go behind the scene creates a new array with double memory alocation always. Ofcorse there are some tricks to append stuff to arrays.

Referencing a slice by its variable name will always reference the entire slice.

A nice trick to append a slice to another slice could be this one. Existing slice will be reasigned (note the asignation) and we just use the append function not on the slice instead we declare it as a new asignment and then use te elipses tis takes the values, and append theam each to the other slice.

```go
package main

import (
    "fmt"
)

func main() {
    mySlice  :=  []int{1,2,3,4,5,6,7}
    fmt.Println(mySlice)

    for _, i := range mySlice {
        fmt.Println("for range loop...", i)
    }

    newSlice:= []int{10,20,30}
    mySlice = append(mySlice, newSlice...)
    fmt.Println(mySlice)
}
```

## Structs

Structs let tu define custom data types, its a way to enhence a data type. ( its like graphql customtyes haha). Structs just like maps need to be created and then can be instanciated and alterd. This must be the introduction to Go  OOP this is modular objects, so it can be considerd. because go lacks of object types adn class type, and they no support inheretance. Go it's not designd to be a OOP.

```go
package main

import {
    "fmt"
}

func main() {
    type courseMeta struct {
        Author string
        Level string
        Rating float64
    }

    somthing := courseMeta{
        Author: "someone",
        Level: "Intermidate",
        Rating: 5,
    }

    fmt.Println(somthing)
}
```