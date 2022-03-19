package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubMult(t *testing.T) {
	t.Parallel()
	type testCase struct {
		fn   func(float64, float64) float64
		a, b float64
		want float64
	}
	tcs := []testCase{
		{fn: calculator.Add, a: 2, b: 2, want: 4},
		{fn: calculator.Add, a: 600, b: 650, want: 1250},
		{fn: calculator.Add, a: -9, b: -10, want: -19},
		{fn: calculator.Add, a: -20, b: 100, want: 80},
		{fn: calculator.Add, a: -7, b: 7, want: 0},
		{fn: calculator.Subtract, a: 10, b: 5, want: 5},
		{fn: calculator.Subtract, a: -100, b: -200, want: 100},
		{fn: calculator.Subtract, a: 999, b: 1000, want: -1},
		{fn: calculator.Multiply, a: 10, b: 100, want: 1000},
		{fn: calculator.Multiply, a: 5, b: -5, want: -25},
		{fn: calculator.Multiply, a: 20, b: 30, want: 600},
	}
	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Add(%.1f, %.1f): want %.1f, got %.1f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var a, b float64
	a = 3
	b = 2
	var want float64 = 1
	got := calculator.Subtract(a, b)

	if want != got {
		t.Errorf("Subtract(%.1f, %.1f): want %.1f, got %.1f", a, b, want, got)
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var a, b float64
	a = 3
	b = 2
	var want float64 = 6
	got := calculator.Multiply(a, b)

	if want != got {
		t.Errorf("Multiply(%.1f, %.1f): want %.1f, got %.1f", a, b, want, got)
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a, b        float64
		want        float64
		errExpected bool
	}{
		{a: 3, b: 3, want: 1, errExpected: false},
		{a: 3, b: 0, want: -999, errExpected: true},
		{a: 10, b: 0, want: -999, errExpected: false},
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%1.f,%1.f): unexpected error status: %v", tc.a, tc.b, err)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%.1f, %.1f): want %.1f, got %.1f", tc.a, tc.b, tc.want, got)
		}

	}

}

func TestSqrt(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a           float64
		want        float64
		errExpected bool
	}{
		{a: 4, want: 2, errExpected: false},
		{a: -3, want: -999, errExpected: true},
		{a: 25, want: 5, errExpected: false},
	}
	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.a)
		errReceived := (err != nil)
		if errReceived != tc.errExpected {
			t.Fatalf("Sqrt(%1.f): unexpected error status: %v", tc.a, err)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Sqrt(%.1f): want %.1f, got %.1f", tc.a, tc.want, got)
		}

	}

}
