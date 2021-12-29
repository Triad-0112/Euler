package main

import "fmt"

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different variables in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}
}
