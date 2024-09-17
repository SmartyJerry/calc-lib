package calc

import "testing"

func TestMultiplication_Calculate(t *testing.T) {
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
		{"Single Zero First", args{5, 0}, 0},
		{"Single Zero Second", args{0, 9}, 0},
		{"Positives", args{5, 3}, 15},
		{"Negative First", args{-2, 5}, -10},
		{"Negative Second", args{2, -14}, -28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplication := Multiplication{}
			got, err := multiplication.Calculate(tt.args.a, tt.args.b)
			assertError(t, nil, err)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
