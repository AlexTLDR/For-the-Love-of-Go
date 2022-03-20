package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Think Like a Programmer: An Introduction to Creative Problem Solving, by V. Anton Spraul"
	got := bookstore.GetBookDetails(catalog, "1")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
