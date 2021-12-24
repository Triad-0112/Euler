package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	a := strings.Join(strings.Fields(id), "")
	b := strings.Split(a, "")
	array := make([]int, len(b))
	total := 0
	for i := range array {
		array[i], _ = strconv.Atoi(b[i])
	}
	for i := 0; i < len(array); i++ {
		if i == 0 || i%2 == 0 {
			array[i] *= 2
			if array[i] > 9 {
				array[i] = array[i] - 9
			}
		}
		total += array[i]
	}
	if total%10 != 0 {
		return false
	} else {
		return true
	}
}
