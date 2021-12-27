package main

import "fmt"

func main() {
	a0 := new(int)
	fmt.Println(a0)
	*a0 = 10
	fmt.Println(*a0)
	x := *a0
	fmt.Println(x)
	b1, b2 := &x, &x
	fmt.Println(b1, b2)

}
