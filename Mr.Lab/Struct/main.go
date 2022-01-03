package main

type Curr struct {
	Currency string
	Name     string
	Country  string
	Number   int
}

/*
EXAMPLE :
func assertEuro(c Curr) bool {
	switch name, curr := "Euro", "EUR"; {
	case c.Name == name:
		return true
	case c.Currency == curr
		return true
	}
	return false
}

This is will result to find any input of string and search it on data set
func find(name string) {
	for i := 0; i<10;i++ {
		c:= currencies[i]
		switch {
			case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name):
			fmt.Println("Found", curr)
		}
	}
}

This is will search for number of interger in dataset
func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("Found", curr)
		}
	}
}
*/
