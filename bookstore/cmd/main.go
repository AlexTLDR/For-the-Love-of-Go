package main

import (
	"bookstore"
	"fmt"
)

func main() {
	var b bookstore.Book
	err := b.SetCategory(bookstore.CategoryScience)
	fmt.Println(err)
}
