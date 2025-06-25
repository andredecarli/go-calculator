package main

type calcImpl struct{}

func NewCalculator() calcImpl {
	return calcImpl{}
}

func (c calcImpl) Add(a, b float64) float64 {
	return a + b
}

func (c calcImpl) Subtract(a, b float64) float64 {
	return a - b
}

func (c calcImpl) Multiply(a, b float64) float64 {
	return a * b
}

func (c calcImpl) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
