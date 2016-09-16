package calc

import (
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		in   string
		want float64
	}{
		{"1+2", 3.0},
		{"1-2", -1.0},
		{"1*2", 2.0},
		{"1/2", .5},
		{"2^3", 8},

		{"2+3*4", 14.0},
		{"2*3+4", 10.0},

		{"FLOOR(2.3)", 2.0},

		{"SIGN(2.3)", 1.0},
		{"SIGN(-2.3)", -1.0},
	}

	for _, c := range cases {
		got := Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
