package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println("\nargsWithProg : \t ", argsWithProg)
	fmt.Println("\nargsWithoutProg : \t ", argsWithoutProg)
	fmt.Println("\narg := os.Args[3] : \t ", arg)
}
