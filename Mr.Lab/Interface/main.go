package main

/*
Go offers the type interface{} or empty interface as a super type.

EXAMPLE:
func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}

*/
