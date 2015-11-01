package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {

	var tests = []struct {
		i, o int
	}{
		{0, 0},
		{1, 1},
		{9, 34},
		{40, 102334155},
		{49, 7778742049},
	}
	for _, c := range tests {
		got := Fibo(c.i)
		if got != c.o {
			t.Errorf("Fibo(%q) == %q, want %q", c.i, got, c.o)
		}
	}

}
