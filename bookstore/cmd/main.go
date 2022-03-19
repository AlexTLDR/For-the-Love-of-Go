package main

import "fmt"

var books []Book

type Book struct {
	author            string
	title             string
	copies            int64
	onSpecial         bool
	specialOffer      string
	royaltyPercentage float64
	priceInCents      int64
}

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(books, book)
	return catalog
}

func main() {

	books = []Book{
		{author: "V. Anton Spraul",
			title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
			copies:            1,
			onSpecial:         false,
			specialOffer:      "0%",
			royaltyPercentage: 15.5,
			priceInCents:      900},

		{author: "John Arundel",
			title:             "For the Love of Go",
			copies:            12,
			onSpecial:         true,
			specialOffer:      "40%",
			royaltyPercentage: 15.5,
			priceInCents:      4500},
	}

	books = AddBook(books, Book{author: "William Kennedy", title: "Ultimate Go Notebook"})

	fmt.Println(books)
	for _, b := range books {
		fmt.Println(b.author)
	}

}
