package main

import "fmt"

type Book struct {
	pages int
}

type Books []Book

func (books Books) Modify() {
	// Modifications on the underlying part of
	// the receiver will be reflected to outside
	// of the method.
	books[0].pages = 500
	// Modifications on the direct part of the
	// receiver will not be reflected to outside
	// of the method.
	books = append(books, Book{789})
}

func main() {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books) // [{500} {456}]
}
