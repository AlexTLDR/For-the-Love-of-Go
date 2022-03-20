package main

import "fmt"

var books []Book

type Book struct {
	Author            string
	Title             string
	Copies            int64
	OnSpecial         bool
	SpecialOffer      string
	RoyaltyPercentage float64
	PriceInCents      int64
}

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(books, book)
	return catalog
}

func main() {

	books = []Book{
		{
			Author:            "V. Anton Spraul",
			Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
			Copies:            1,
			OnSpecial:         false,
			SpecialOffer:      "0%",
			RoyaltyPercentage: 15.5,
			PriceInCents:      900,
		},

		{
			Author:            "John Arundel",
			Title:             "For the Love of Go",
			Copies:            12,
			OnSpecial:         true,
			SpecialOffer:      "40%",
			RoyaltyPercentage: 15.5,
			PriceInCents:      4500,
		},
	}

	books = AddBook(books, Book{Author: "William Kennedy", Title: "Ultimate Go Notebook"})

	fmt.Println(books)
	for _, b := range books {
		fmt.Println(b.Author)
	}

}
