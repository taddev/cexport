package cexport

import "fmt"

/*
#include <stdio.h>
extern void ACFunction();
*/
import "C"

//export AGoFunction
func AGoFunction() {
	fmt.Println("Hello, GO!")
}

func RunCGOC() {
	C.ACFunction()
}
