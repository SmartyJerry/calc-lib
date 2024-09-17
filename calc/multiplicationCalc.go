package calc

type Multiplication struct{}

func (multiplication Multiplication) Calculate(a, b int) (int, error) { return a * b, nil }
