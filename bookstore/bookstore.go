package bookstore

import "fmt"

type Book struct {
	ID                string
	Author            string
	Title             string
	Copies            int64
	OnSpecial         bool
	SpecialOffer      string
	RoyaltyPercentage float64
	PriceInCents      int64
}

type Catalog map[string]Book

func GetAllBooks(catalog Catalog) []Book {
	books := []Book{}
	for _, b := range catalog {
		books = append(books, b)
	}
	return books
}

func GetBookDetails(catalog Catalog, ID string) string {
	b := catalog[ID]
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
}
