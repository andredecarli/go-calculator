package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibCalculator_Sum(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a + b

	res := calc.Add(a, b)

	assert.Equal(t, want, res)
}

func TestLibCalculator_Subtract(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a - b

	res := calc.Subtract(a, b)

	assert.Equal(t, want, res)
}

func TestLibCalculator_Multiply(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a * b

	res := calc.Multiply(a, b)

	assert.Equal(t, want, res)
}

func TestLibCalculator_Divide(t *testing.T) {
	t.Run("should divide correctly", func(t *testing.T) {
		calc := NewCalculator()
		a := 1.0
		b := 2.0
		want := a / b

		res, err := calc.Divide(a, b)

		assert.NoError(t, err)
		assert.Equal(t, want, res)
	})

	t.Run("should return an error when dividing by zero", func(t *testing.T) {
		calc := NewCalculator()
		a := 1.0
		b := 0.0

		_, err := calc.Divide(a, b)

		assert.Error(t, err)
		assert.Equal(t, ErrDivideByZero, err)
	})
}
