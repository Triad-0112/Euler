package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	string := ("7160 6873 781")
	a := strings.Join(strings.Fields(string), "")
	b := strings.Split(a, "")
	var c bool
	array := make([]int, len(b))
	total := 0
	for i := range array {
		array[i], _ = strconv.Atoi(b[i])
	}
	fmt.Println(array)
	for i := len(array) - 1; i >= 0; i-- {
		if len(array)%2 == 0 {
			if i%2 == 0 || i == 0 {
				array[i] *= 2
				if array[i] > 9 {
					array[i] = array[i] - 9
				}
			}
		} else {
			if i == len(array)-2 || (i-1)%2 == 0 {
				array[i] *= 2
				if array[i] > 9 {
					array[i] = array[i] - 9
				}
			}
		}
		total += array[i]
	}
	if total%10 != 0 {
		c = false
	} else {
		c = true
	}
	fmt.Println(array)
	fmt.Println(total)
	fmt.Println(c)
}
