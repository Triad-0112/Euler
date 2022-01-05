package main

import "fmt"

type DoubleValue [2]int
type TripleValue [3]int
type FourValue [4]int

func Encrypt2(a DoubleValue) int {
	var Encrypted int
	for i := 0; i <= len(a)-1; i++ {
		Encrypted += a[i]
	}
	Encrypted += 5
	Encrypted *= 3
	return Encrypted
}
func Encrypt3(a TripleValue) int {
	var Encrypted int
	for i := 0; i <= len(a)-1; i++ {
		Encrypted += a[i]
	}
	Encrypted += 7
	Encrypted *= 2
	return Encrypted
}
func Encrypt4(a FourValue) int {
	var Encrypted int
	for i := 0; i <= len(a)-1; i++ {
		Encrypted += a[i]
	}
	Encrypted -= 150
	if Encrypted < 0 {
		Encrypted *= -50
	} else {
		Encrypted *= 2
	}
	return Encrypted
}
func main() {
	var a DoubleValue
	a[0] = 5
	a[1] = 10
	fmt.Printf("Encrypt Number for %d is %d\n\n", a, Encrypt2(a))
	var b TripleValue
	b[0] = Encrypt2(a)
	b[1] = 100
	b[2] = 30
	fmt.Printf("Encrypt Number for %d is %d\n\n", b, Encrypt3(b))
	var c FourValue
	c[0] = Encrypt2(a)
	c[1] = Encrypt3(b)
	c[2] = Encrypt3(b) - Encrypt2(a)
	c[3] = Encrypt2(a) * 3
	fmt.Printf("Encrypt Nnumber for %d is %d\n\n", c, Encrypt4(c))

}
