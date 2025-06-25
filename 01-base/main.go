package main

import "fmt"

const (
	ADICAO        int = iota + 1 // 1
	SUBTRACAO                    // 2
	MULTIPLICACAO                // 3
	DIVISAO                      // 4
)

func main() {
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
		result = firstNumber + secondNumber
	case SUBTRACAO:
		result = firstNumber - secondNumber
	case MULTIPLICACAO:
		result = firstNumber * secondNumber
	case DIVISAO:
		if secondNumber == 0 {
			fmt.Println("Impossível dividir por zero.")
			return
		}
		result = firstNumber / secondNumber
	}

	fmt.Printf("Resultado: %.2f\n", result)
}
