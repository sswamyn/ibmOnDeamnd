package main

import (
	"flag"
	"fmt"
	"os"
)

// Simple error checker to help with multiple occurrence of err checking.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Usage pgmName -dataFile=VSR1PR01_HIST.VSR1PR01.out -offset=40975340  -dataLength=14140
	datafilePtr := flag.String("dataFile", "DataFile.out", "The data file that needs to be parsed. \n\t -dataFile=VSR1PR01_HIST.VSR1PR01.out")
	offsetPtr := flag.Int("offset", 0, "The offset from the begin. [all positive integers including zero]. \n\t -offset=40975340")
	lengthPtr := flag.Int("dataLength", 1, "The length of data segment [a positive integers]. \n\t -dataLength=14140")

	/**
	It’s also possible to declare an option that uses an existing var declared elsewhere in the program.
	Note that we need to pass in a pointer to the flag declaration function.
	*/
	// Now that pointer for the command line arugments are defines, let us parse the arguments
	flag.Parse()

	// Let us print the arguments back to the user
	fmt.Println("Data file : \t", *datafilePtr)
	fmt.Println("Offset : \t", *offsetPtr)
	fmt.Println("Length : \t", *lengthPtr)
	// Tail to capture any straggler arguments
	fmt.Println("straggler arguments : \t", flag.Args())

	output := processFile(*datafilePtr, *offsetPtr, *lengthPtr)

	fmt.Printf("%s\n", "--------------------------------------------------------------------------------------------------------")
	fmt.Printf("%s ", output)
	fmt.Printf("\n%s", "--------------------------------------------------------------------------------------------------------")
	print("\n Goodbye")
}

func processFile(dataFile string, offset int, length int) (extractedData string) {

	// If you do not want the whole file in memory, just open .. and read a few bytes [b]
	// f, err := os.Open("C:\\Users\\pc23sxs\\Downloads\\IBM - OnDemand-datasample 2023-01-06\\data\\test2.txt")
	f, err := os.Open(dataFile)
	check(err)

	fs, err := f.Stat()
	check(err)
	var size = fs.Size()

	var checkLength = offset + length
	if int64(checkLength) > size {
		print("oops! some thing is not adding up! :) \n did you get it 'add' :) ")
		print("\n\t *********** File size is %d, offset(%d) plus length(%d) is %d ***********\n", size, offset, length, checkLength)
		os.Exit(2)
	}

	//o2, err := f.Seek(40975340, 0)
	o2, err := f.Seek(int64(offset), 0)
	check(err)
	//b2 := make([]byte, 14140)
	b2 := make([]byte, length)
	n2, err := f.Read(b2)
	check(err)
	// Info print - to show the actual parameters the program is using
	fmt.Printf("%s's size is:  %d : \n", dataFile, size)
	fmt.Printf("%d bytes @ %d: \n", n2, o2)

	f.Close()

	return string(b2[:n2])
}
