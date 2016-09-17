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

		{"0.5 * 0.5", .25},

		{"2+3*4", 14.0},
		{"2*3+4", 10.0},

		{"FLOOR(2.3)", 2.0},

		{"ROUND(2.3)", 2.0},
		{"ROUND(2.6)", 3.0},
		{"ROUND(-2.3)", -2.0},
		{"ROUND(-2.6)", -3.0},

		{"SIGN(2.3)", 1.0},
		{"SIGN(-2.3)", -1.0},

		{"SIGN(FLOOR((-5+0.999999)/2))", -1.0},
		{"SIGN(FLOOR((5+0.999999)/2))", 1.0},
	}

	for _, c := range cases {
		got := Solve(c.in)
		if got != c.want {
			t.Errorf("Solve(%q) == %f, want %f", c.in, got, c.want)
		}
	}
}
