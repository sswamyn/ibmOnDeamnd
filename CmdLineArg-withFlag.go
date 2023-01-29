package main

import (
	"flag"
	"fmt"
)

func main() {

	datafilePtr := flag.String("dataFile", "DataFile.out", "The data file that needs to be parsed.")
	offsetPtr := flag.Int("Offset", 0, "The offset from the begin. [all positive integers including zero].")
	lengthPtr := flag.Int("dataLength", 1, "The length of data segment [a positive integers].")

	/**
	It’s also possible to declare an option that uses an existing var declared elsewhere in the program.
	Note that we need to pass in a pointer to the flag declaration function.
	*/
	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")

	// Now that pointer for the command line arugments are defines, let us parse the arguments
	flag.Parse()

	// Let us print the arguments back to the user
	fmt.Println("Data file : \t", *datafilePtr)
	fmt.Println("Offset : \t", *offsetPtr)
	fmt.Println("Length : \t", *lengthPtr)
	//fmt.Println("svar:", svar)
	// Tail to capture any straggler arguments
	fmt.Println("straggler arguments : \t", flag.Args())

}
