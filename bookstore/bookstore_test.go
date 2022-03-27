package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var catalog = bookstore.Catalog{
	"1": {
		ID:                "1",
		Author:            "V. Anton Spraul",
		Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
		Copies:            1,
		OnSpecial:         false,
		SpecialOffer:      "0%",
		RoyaltyPercentage: 15.5,
		PriceInCents:      900,
	},

	"2": {
		ID:                "2",
		Author:            "John Arundel",
		Title:             "For the Love of Go",
		Copies:            12,
		OnSpecial:         true,
		SpecialOffer:      "40%",
		RoyaltyPercentage: 15.5,
		PriceInCents:      4500,
	},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Author:            "V. Anton Spraul",
		Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
		Copies:            1,
		OnSpecial:         false,
		SpecialOffer:      "0%",
		RoyaltyPercentage: 15.5,
		PriceInCents:      900,
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	want := []bookstore.Book{
		{
			ID:                "1",
			Author:            "V. Anton Spraul",
			Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
			Copies:            1,
			OnSpecial:         false,
			SpecialOffer:      "0%",
			RoyaltyPercentage: 15.5,
			PriceInCents:      900,
		},

		{
			ID:                "2",
			Author:            "John Arundel",
			Title:             "For the Love of Go",
			Copies:            12,
			OnSpecial:         true,
			SpecialOffer:      "40%",
			RoyaltyPercentage: 15.5,
			PriceInCents:      4500,
		},
	}
	got := catalog.GetAllBooks()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	b := bookstore.Book{
		ID:                "1",
		Author:            "V. Anton Spraul",
		Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
		Copies:            1,
		OnSpecial:         false,
		SpecialOffer:      "0%",
		RoyaltyPercentage: 15.5,
		PriceInCents:      900,
	}
	want := []bookstore.Book{
		b,
	}

	catalog.AddBook(b)

	got := catalog.GetAllBooks()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetCategory(t *testing.T) {
	b := bookstore.Book{
		Title: "Think Like a Programmer: An Introduction to Creative Problem Solving",
	}
	err := b.SetCategory(-1)
	if err == nil {
		t.Error("want error setting foo category, got nil")
	}
	err = b.SetCategory(bookstore.CategoryScience)
	if err != nil {
		t.Errorf("want no error setting valid category, got %v", err)
	}
	got := b.Category()
	want := bookstore.CategoryScience
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetDiscountPercent(t *testing.T) {
	b := bookstore.Book{
		Title: "Think Like a Programmer: An Introduction to Creative Problem Solving",
	}
	err := b.SetDiscountPercent(101)
	if err == nil {
		t.Error("want error setting discount percent < 0 or > 0, got nil")
	}
	err = b.SetDiscountPercent(25)
	if err != nil {
		t.Errorf("want no error setting valid category, got %d", err)
	}
	got := b.DiscountPercent()
	want := 25
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Think Like a Programmer: An Introduction to Creative Problem Solving, by V. Anton Spraul"
	got := catalog.GetBookDetails("1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNetPrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:        "The Go Programming Language",
		Author:       "Alan A. A. Donovan",
		PriceInCents: 100,
	}
	want := 75
	got := b.NetPrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSalePrice(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		ID:                "1",
		Author:            "V. Anton Spraul",
		Title:             "Think Like a Programmer: An Introduction to Creative Problem Solving",
		Copies:            1,
		OnSpecial:         false,
		SpecialOffer:      "0%",
		RoyaltyPercentage: 15.5,
		PriceInCents:      900,
	}
	want := 450
	got := b.SalePrice()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetTitle(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "Black Hat Go",
	}
	b.SetTitle("White Hat Go")
	want := "White Hat Go"
	got := b.Title
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:        "Black Hat Go",
		PriceInCents: 1000,
	}
	b.SetPriceCents(1500)
	want := 1500
	got := b.PriceInCents
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
