package calc

type Addition struct{}

func (addition Addition) Calculate(a, b int) (int, error) { return a + b, nil }
