package calc

import "testing"

func TestSubtraction_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zeros", args{0, 0}, 0},
		{"Positives", args{5, 1}, 4},
		{"Negative First", args{-2, 5}, -7},
		{"Negative Second", args{2, -5}, 7},
		{"Larger Second", args{2, 5}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subtraction := Subtraction{}
			got, err := subtraction.Calculate(tt.args.a, tt.args.b)
			assertError(t, nil, err)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
