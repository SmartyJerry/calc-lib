package calc

type Subtraction struct{}

func (subtraction Subtraction) Calculate(a, b int) (int, error) { return a - b, nil }
