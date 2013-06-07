package cexport

import "fmt"
import "C"

//export AGoFunction
func AGoFunction() {
	fmt.Println("Hello, GO!")
}

