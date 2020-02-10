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

## Error handling

There are many ways to handle errors in Go, but the main premise its that it doesnt exist the try, catch block, or somthing like that. So its necesary to do the **nil** trick. all functions return two values, the variable and an error, if there the error has an object that its not nil then the variable was pointed to an object, so an error.

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}

type error interface {
    Error() string
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

sometimes people add an interface or a struct to that so that the error when it pops outyou can react in the same way

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

Go only uses for loops of three types: infinite (while), range, conditional. you can use the ```break``` and  ```continue``` inside a loop, these givs life to the loop. you cna break stuff within a conditional, and stuff like that.

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

"Named lists containing elements of the same type". Arrays fot fixed lenght when created (cos memory alocation), so for dinamic ranges you need to use Slices, The thing is Slices in essence are arrays managed by Go, when you declare a Slice in reality Go creates an array. These meaning, that slices are always representation of memory alogation, so if you change the data on a slice will **ALWAYS** be reflected on the array. What Go does is creating a Slice inside an array, that when it reach the limit it doubles the size of the array, for more reference run the funcion of array expansión. creating arrays need the ```make()``` function in order to make new slices.

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

## Maps
maps are the way to organize the data like a traditional hash maps arrays, and the asaignment are like all the lenguajes, all the access of the stuff are the same with bracket notation, like always investigate all the methos that comes with a data structure like maps, and one interesting function is the *delete()* the first parameter is the map the second one the key.
All the data types of the maps have to be the same type, thats one of the **homogenius** caracteristic of Go

```go
m :=map[string]int{"key":"value"}
fmt.Println(m)
fmt.Println(m["foo"])

m["foo"] = "somthing else"
delete(m, "foo")
```

## Structs

Structs let tu define custom data types, its a way to enhence a data type. ( its like graphql customtyes haha). Structs just like maps need to be created and then can be instanciated and alterd. This must be the introduction to Go  OOP this is modular objects, so it can be considerd. because go lacks of object types adn class type, and they no support inheretance. Go it's not designd to be a OOP, but these can help to create a ilution, strucst always are user defined.
The declaration of values of structs are like the map declaration.

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

## looping trough a map

these is the usual way to loop trought the first value to get all the keys, the problem in order to get the key and vlue you need to declare a read only variable with _, because if you declare to variables to access a map and dont use the first it will throw an error of not using a declared variable.
```go
wellKnownPorts:= map[string]int{"http": 80, "https": 443}
for k := range wellKnownPorts {
    fmt.Println(k)
}

wellKnownPorts:= map[string]int{"http": 80, "https": 443}
for _, k := range wellKnownPorts {
    fmt.Println(k)
}
```

# Panic state

its a situation similar to exeptions in other lenguajes but it sends a signal to all the aplication so it can strop and notify the state to all the parts, panic its a panic object created with the *panic()* method these tell you where exactl the panic happend.
Recover from panic its poisbile and necesry

