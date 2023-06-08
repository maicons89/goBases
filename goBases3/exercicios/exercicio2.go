//Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
//qual possamos passar N quantidade de números inteiros e devolva a média.
//Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro

package main

import (
	"errors"
	"fmt"
)

func calcularMedia(notas ...int) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("Nenhuma nota foi fornecida")
	}

	soma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Nota negativa encontrada")
		}
		soma += nota
	}

	media := float64(soma) / float64(len(notas))
	return media, nil
}

func main() {
	notas := []int{10, 8, 7, 9}

	media, err := calcularMedia(notas...)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("A média das notas é: %.2f\n", media)
}
