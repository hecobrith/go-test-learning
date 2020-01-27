package main
import (
	"fmt"
	"runtime"
	"reflect"
)
// to declare variables are necesary to put them inside the var statement ourside the main func
// declaring a variable without a value will always become 0 allocated in memory
// these variables a global in scope
var (
	name string
	course string
	module float64
)

// Getting the runtime.GOOS OS
func main() {
	fmt.Println("hello from", runtime.GOOS )
	fmt.Println( reflect.TypeOf(name) )
	fmt.Println( reflect.TypeOf(module) )
}