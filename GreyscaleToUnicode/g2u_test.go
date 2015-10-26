package main

import (
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		s    int
		want string
	}{
		{30, "d"},
		{21, "W"},
		{33, "w"},
	}
	for _, c := range tests {
		got := Itou(c.s)
		if got != c.want {
			t.Errorf("Itou(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
