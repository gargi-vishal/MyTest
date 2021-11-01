package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	// Print the name of the command that invoked this
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}

