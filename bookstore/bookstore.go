package bookstore

import "fmt"

type Book struct {
	ID                string
	Author            string
	Title             string
	Copies            int
	OnSpecial         bool
	SpecialOffer      string
	RoyaltyPercentage float64
	PriceInCents      int
	discountPercent   int
	category          string
}

type Catalog map[string]Book

func (catalog Catalog) GetAllBooks() []Book {
	books := []Book{}
	for _, b := range catalog {
		books = append(books, b)
	}
	return books
}

func (catalog *Catalog) AddBook(b Book) {
	(*catalog)[b.ID] = b
}

func (catalog Catalog) GetBookDetails(ID string) string {
	b := catalog[ID]
	return fmt.Sprintf("%s, by %s", b.Title, b.Author)
}

func (b Book) NetPrice() int {
	saving := b.PriceInCents * b.discountPercent / 100
	return b.PriceInCents - saving
}

func (b Book) SalePrice() int {
	/*sale := b.PriceInCents * b.HalfPrice / 100
	return b.PriceInCents - sale*/
	return b.PriceInCents / 2
}

func (b *Book) SetTitle(t string) {
	b.Title = t
}

func (b *Book) SetPriceCents(i int) {
	b.PriceInCents = i
}

func (b *Book) SetCategory(c string) error {
	if c != "Science" {
		return fmt.Errorf("invalid category %q", c)
	}
	b.category = c
	return nil
}

func (b Book) Category() string {
	return b.category
}

func (b *Book) SetDiscountPercent(d int) error {
	if d < 0 || d > 100 {
		return fmt.Errorf("invalid discount %d", d)
	}
	b.discountPercent = d
	return nil
}

func (b Book) DiscountPercent() int {
	return b.discountPercent
}
