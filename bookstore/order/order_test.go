package order_test

import (
	"bookstore/order"
	"testing"
)

func TestNewValid(t *testing.T) {
	t.Parallel()
	want := "1"
	order, err := order.New(want)
	if err != nil {
		t.Error(err)
	}
	got := order.ID()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()
	want := ""
	_, err := order.New(want)
	if err == nil {
		t.Error("wanted error for invalid ID, got nil")
	}
}

func TestSetID(t *testing.T) {
	t.Parallel()
	want := "1"
	order, err := order.New(want)
	if err != nil {
		t.Error(err)
	}
	err = order.SetID("")
	if err == nil {
		t.Error("wanted error on setting invalid ID, got nil")
	}
	want = "don't want an empty string"
	err = order.SetID(want)
	if err != nil {
		t.Error(err)
	}
	got := order.ID()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
