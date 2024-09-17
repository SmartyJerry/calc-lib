package calc

import (
	"errors"
	"testing"
)

func assertError(t *testing.T, expected, actual error) {
	if !errors.Is(actual, expected) {
		t.Helper()
		t.Errorf("\n"+
			"expected: [%v]\n"+
			"actual: [%v]", expected, actual)
	}
}

func TestDivision_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{"Zeros", args{0, 0}, 0, errDivideByZero},
		{"Positives", args{10, 5}, 2, nil},
		{"Negative First", args{-20, 5}, -4, nil},
		{"Negative Second", args{28, -7}, -4, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			division := Division{}
			got, err := division.Calculate(tt.args.a, tt.args.b)
			assertError(t, tt.err, err)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
