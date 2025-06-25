package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("impossivel dividir por zero")
var ErrSqrtOfNegative = errors.New("impossivel tirar raíz quadrada de número negativo")

const (
	Add      int = iota + 1 // 1
	Subtract                // 2
	Multiply                // 3
	Divide                  // 4
	Power                   // 5
	Sqrt                    // 6
)

func main() {
	calc := ScientificCalculator{}

	fmt.Println("Go-Calculator")

	var firstNumber float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanf("%f", &firstNumber)

	var operation int
	fmt.Print("Digite a operação:\n1-Adição, 2-Subtração\n3-Multiplicação, 4-Divisão\n5-Potencia, 6-Raiz Quadrada\n> ")
	fmt.Scanf("%d", &operation)

	for operation < 1 || operation > 6 {
		fmt.Println("Operação inválida!")

		fmt.Print("Digite a operação:\n1-Adição, 2-Subtração\n3-Multiplicação, 4-Divisão\n5-Potencia, 6-Raiz Quadrada\n> ")
		fmt.Scanf("%d", &operation)
	}

	var secondNumber float64
	if operation != Sqrt {
		fmt.Print("Digite o segundo número: ")
		fmt.Scanf("%f", &secondNumber)
	}

	var result float64
	switch operation {
	case Add:
		result = calc.Add(firstNumber, secondNumber)
	case Subtract:
		result = calc.Subtract(firstNumber, secondNumber)
	case Multiply:
		result = calc.Multiply(firstNumber, secondNumber)
	case Divide:
		quotient, err := calc.Divide(firstNumber, secondNumber)
		if err != nil {
			fmt.Println("Erro na divisão:", err.Error())
			return
		}
		result = quotient
	case Power:
		result = calc.Power(firstNumber, secondNumber)
	case Sqrt:
		res, err := calc.Sqrt(firstNumber)
		if err != nil {
			fmt.Println("Erro na raiz quadrada:", err.Error())
			return
		}
		result = res
	}

	fmt.Printf("Resultado: %.2f\n", result)
}
