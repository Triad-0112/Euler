package hamming

import "errors"

func Distance(a, b string) (int, error) {
	space := 0
	if len(a) != len(b) {
		return -1, errors.New("need inputs with same length")
	}
	for i := range a {
		if b[i] != a[i] {
			space++
		}
	}
	return space, nil
}
