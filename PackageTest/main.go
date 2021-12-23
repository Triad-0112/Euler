package main

import (
	"fmt"

	Damage "github.com/Triad-0112/GoCode/TestPackage/Idonk"
)

func main() {
	a := 1000
	b := 50000
	Damage.Attack(a, b)
	fmt.Println(b)
}
