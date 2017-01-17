package factorial

import "testing"

type testpair struct {
	input    int
	expected int
}

func TestFactorial(t *testing.T) {
	var tests = []testpair{
		{-1, 1},
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{5, 120},
		{10, 3628800},
	}
	for _, pair := range tests {
		result := Factorial(pair.input)
		if result != pair.expected {
			t.Error("Factorial(", pair.input, ") = ", result, " - expected ",
				pair.expected)
		}
	}
}
