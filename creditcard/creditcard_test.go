package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewValid(t *testing.T) {
	t.Parallel()
	want := "12345678"
	card, err := creditcard.New(want)
	if err != nil {
		t.Error(err)
	}
	got := card.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()
	want := ""
	_, err := creditcard.New(want)
	if err == nil {
		t.Error("want error from invalid number, got nil")
	}
}

func TestSetNumber(t *testing.T) {
	t.Parallel()
	want := "12345678"
	card, err := creditcard.New(want)
	if err != nil {
		t.Error(err)
	}
	err = card.SetNumber("")
	if err == nil {
		t.Error("want error on setting invalid number, got nil")
	}
	want = "anything other than empty string"
	err = card.SetNumber(want)
	if err != nil {
		t.Error(err)
	}
	got := card.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
