//Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
//notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
//de suas notas.
//Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
//máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
//definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
//indicado na função anterior

package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type CalculationFunc func(...int) (int, error)

func operation(calculationType string) (CalculationFunc, error) {
	switch calculationType {
	case minimum:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("Nenhuma nota foi fornecida")
			}

			min := numbers[0]
			for _, num := range numbers {
				if num < min {
					min = num
				}
			}

			return min, nil
		}, nil

	case average:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("Nenhuma nota foi fornecida")
			}

			sum := 0
			for _, num := range numbers {
				sum += num
			}

			average := sum / len(numbers)
			return average, nil
		}, nil

	case maximum:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("Nenhuma nota foi fornecida")
			}

			max := numbers[0]
			for _, num := range numbers {
				if num > max {
					max = num
				}
			}

			return max, nil
		}, nil

	default:
		return nil, errors.New("Tipo de cálculo inválido")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	avgFunc, err := operation(average)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	numbers := []int{2, 3, 3, 4, 10, 2, 4, 5}

	minValue, err := minFunc(numbers...)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	avgValue, err := avgFunc(numbers...)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	maxValue, err := maxFunc(numbers...)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Valor mínimo: %d\n", minValue)
	fmt.Printf("Valor médio: %d\n", avgValue)
	fmt.Printf("Valor máximo: %d\n", maxValue)
}
