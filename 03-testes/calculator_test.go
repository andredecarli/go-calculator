package main

import "testing"

func TestCalculator_Sum(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a + b
	res := calc.Add(a, b)

	if res != want {
		t.Errorf(`Sum(1.0, 2.0) = %f, want match for %f`, res, want)
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a - b
	res := calc.Subtract(a, b)

	if res != want {
		t.Errorf(`Subtract(1.0, 2.0) = %f, want match for %f`, res, want)
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := NewCalculator()
	a := 1.0
	b := 2.0
	want := a * b
	res := calc.Multiply(a, b)

	if res != want {
		t.Errorf(`Multiply(1.0, 2.0) = %f, want match for %f`, res, want)
	}
}

func TestCalculator_Divide(t *testing.T) {
	t.Run("should divide correctly", func(t *testing.T) {
		calc := NewCalculator()
		a := 1.0
		b := 2.0
		want := a / b

		res, err := calc.Divide(a, b)

		if res != want || err != nil {
			t.Errorf(`Divide(1.0, 2.0) = %f, %v, want match for %f, nil`, res, err, want)
		}
	})

	t.Run("should return an error when dividing by zero", func(t *testing.T) {
		calc := NewCalculator()
		a := 1.0
		b := 0.0
		want := 0.0

		res, err := calc.Divide(a, b)

		if res != want || err == nil {
			t.Errorf(`Divide(1.0, 0.0) = %f, %v, want match for %f, %v`, res, err, want, ErrDivideByZero)
		}
	})
}
