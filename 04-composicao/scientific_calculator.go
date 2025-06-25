package main

import "math"

type ScientificCalculator struct {
	BaseCalculator
}

func (c ScientificCalculator) Power(a, b float64) float64 {
	return math.Pow(a, b)
}

func (c ScientificCalculator) Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrSqrtOfNegative
	}
	return math.Sqrt(a), nil
}
