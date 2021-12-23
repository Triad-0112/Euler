package main

import (
	"fmt"
	"math"
)

func Primes(n int) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println(i)
			n /= i
			i--
		}
	}
	if n > 0 {
		fmt.Println(n)
	}
}
func main() {
	number := 600851475143
	Primes(number)
}
