package main

import (
	"flag"
	"fmt"
	//"os"
	//"strings"
)

func main() {
	name := flag.String("n", "", "Nama")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)

	/*var nameArg string
	for index, arg := range os.Args {
		alt := "-n="
		x := strings.Index(arg, alt)
		if x > -1 {
			nameArg = arg[x+len(alt):]
			continue
		}
		if arg == "-n" || arg == "--n" {
			nameArg = os.Args[index+1]
			continue
		}
	}
	fmt.Printf("Hello, %s!\n", nameArg) */
}
