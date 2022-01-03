package main

import "fmt"

/*
type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}
*/

/*

For condition is same as while do loop in any language that have it,
Example :
for x < 10 {
	...
}

There is something called loop for, it will loop over until the funcion is breaking by typing "break" :
for {
	...
}

Traditional way to loop, it will loop until it met requirement of it. It will stop when x >= len(something):
for x:=0, x<len(something);x++{
	...
}

FOr range, you can call this indexing. Its was special in go:
for i, val := range values {
	...
}

You can also do only logical loop like:
for i < len(something) {
	...
}
*/

/*
RANGE SPECIAL
Loop over array or slice:
for i, v := range dataset {
	...
}
Explain : Range producing two values, where i is the loop index and v is the value v[i] from the collection.



Loop over string value:
for i, v := range "Hello" {
	...
}
Explain : Range producing two values, where i is the index of byte in the string v is value utf-8 encoded byte at v[i] returned as rune.




Loop over map:
for k, v := range map[k]v {
	...
}
Explain : Range producing two values, where k is assigned the value of the map key of type K and v get stored  map[k] of type v.



Loop on channel:
for c := range ch {
	...
}
Explain : Assigns each value received from the channel to variable c with each iteration
*/

// EXAMPLE of idexing
var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}

func search(w string) int {
	count := 0
	//DoSearch:
	for i := range words { // same as for i := 0; i<len(words)...
		for k := range words[i] { // same as for i := 0; i<len(words[i])
			if words[i][k] == w {
				count += 1
				fmt.Printf("Found %s!!, total %d\n", w, count)
				//break DoSearch
			}
		}
	}
	return count
}
func main() {
	search("fix")
	search("break")
	search("right")
}
