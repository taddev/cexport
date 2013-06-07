package cexport

/*
#include <stdio.h>

extern void AGoFunction();

void ACFunction() {
    printf("Hello, C!\n");
    AGoFunction();
}
*/
import "C"

func Example() {
	C.ACFunction()
}