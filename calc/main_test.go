package calc_test

import (
	"my-app/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 2, rhs: 1, want: 3},
	}
	for _, test := range tests {
		test := test // Create a local variable and assign the value of test to it
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := calc.Add(test.lhs, test.rhs)
			if got != test.want {
				t.Errorf("%v: want %v, but %v", test.name, test.want, got)
			}
		})
	}
}
