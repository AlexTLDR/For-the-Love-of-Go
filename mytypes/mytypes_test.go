package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Parallel()
	input := mytypes.MultiString{
		"a",
		"b",
		"c",
	}
	want := "abc"
	got := input.Join()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestUppercase(t *testing.T) {
	t.Parallel()
	var input mytypes.StringUppercaser
	input.WriteString("hello")
	want := "HELLO"
	got := input.Uppercase()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSum(t *testing.T) {
	t.Parallel()
	var input mytypes.IntBuilder
	input.Contents = []int{1, 2, 3}
	want := 6
	got := input.Sum()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
