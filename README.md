# Exporting both C and Go Functions #

## Overview ##
The first thing to note is that exported functions in Go must start with a capital letter. The other piece of note new to Go1 is that you can't export a Go function to C and use it in C in the same file. 

Based on code examples from the [go-wiki][3].

From [cgo documentation][1]

> Using //export in a file places a restriction on the preamble: since it is copied into two different C output files, it must not contain any definitions, only declarations. Definitions must be placed in preambles in other files, or in C source files.

My library is available on [Github.com][2].

## cexport Library ##
This library is defined by two files. One holds the C code and exports a Go function for Go use. The second file exports a Go function for C use. To use this library simply import it to a *main* package and call `cexport.Example()`. This will invoke the C function which prints a statement then calls the Go function from C. The method is overly complex but is used to show how these processes are handled.

## Using cexport Library ##
This example program knows nothing about the C code but has complete access to it through the wrapper function `cexport.Example()` we created earlier.

To download this library for your own use do the following, assuming you have setup your Go working directory correctly.

go get github.com/taddevries/cexport

Then write the following program in a your own Go src directory.

### prog1.go ###

    package main

    import "github.com/taddevries/cexport"

    func main() {
        cexport.Example()
    }

### Output ###
Run the program like any other Go program and you're pinning for the fjords......

    > go run prog1.go
    Hello, C!
    Hello, Go!

[1]: http://golang.org/cmd/cgo/ "Command cgo"
[2]: https://github.com/taddevries/cexport
[3]: https://code.google.com/p/go-wiki/wiki/cgo


<!-- rBqBkZXgqodTVH6ncXY5 -->
