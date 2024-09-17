package calc

import "errors"

type Division struct{}

func (division Division) Calculate(a, b int) (int, error) {
	if b == 0 {
		return 0, errDivideByZero
	}
	return a / b, nil
}

var (
	errDivideByZero = errors.New("cannot divide by zero")
)
