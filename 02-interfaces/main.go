package main

import (
	"errors"
	"fmt"
)

type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}

var ErrDivideByZero = errors.New("impossivel dividir por zero")

const (
	ADICAO        int = iota + 1 // 1
	SUBTRACAO                    // 2
	MULTIPLICACAO                // 3
	DIVISAO                      // 4
)

func main() {
	var calc Calculator = NewCalculator()

	fmt.Println("Go-Calculator")

	var firstNumber float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%f", &firstNumber)

	var operation int
	fmt.Print("Digite a operação: 1-Adição, 2-Subtração, 3-Multiplicação, 4-Divisão: ")
	fmt.Scanf("%d", &operation)

	for operation < 1 || operation > 4 {
		fmt.Println("Operação inválida!")

		fmt.Print("Digite a operação: 1-Adição, 2-Subtração, 3-Multiplicação, 4-Divisão: ")
		fmt.Scanf("%d", &operation)
	}

	var secondNumber float64
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%f", &secondNumber)

	var result float64
	switch operation {
	case ADICAO:
		result = calc.Add(firstNumber, secondNumber)
	case SUBTRACAO:
		result = calc.Subtract(firstNumber, secondNumber)
	case MULTIPLICACAO:
		result = calc.Multiply(firstNumber, secondNumber)
	case DIVISAO:
		quotient, err := calc.Divide(firstNumber, secondNumber)
		if err != nil {
			fmt.Println("Erro na divisão:", err.Error())
			return
		}
		result = quotient
	}

	fmt.Printf("Resultado: %.2f\n", result)
}
