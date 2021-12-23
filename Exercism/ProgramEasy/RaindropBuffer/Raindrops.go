//Package raindrops takes an integer value and "translates" it to a particular raindrop sound by determining the integer's prime factorization.
package raindrops

import (
	"bytes"
	"strconv"
)

var rules = []struct {
	Num  int
	Word string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(num int) string {
	var b bytes.Buffer
	for _, rule := range rules {
		if num%rule.Num == 0 {
			b.WriteString(rule.Word)
		}
	}
	if b.String() == "" {
		b.WriteString(strconv.Itoa(num))
	}
	return b.String()
}
