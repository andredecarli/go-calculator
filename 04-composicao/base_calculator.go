package main

type BaseCalculator struct{}

func (c BaseCalculator) Add(a, b float64) float64 {
	return a + b
}

func (c BaseCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c BaseCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c BaseCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
