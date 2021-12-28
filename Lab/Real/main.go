package main

import (
	"fmt"
)

func main() {
	var myNum float64 = 80010000000000000000000
	myNumStr := fmt.Sprintf("%.f", myNum)
	expected := "250000000000000000000000000"
	fmt.Printf("expected: %v\n", expected)
	fmt.Printf("got:  %v\n%T\nGot:  %v\n%T", myNumStr, myNumStr, myNum, myNum)
}
