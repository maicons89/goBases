package main

import (
	"fmt"
)

func main() {

	var idade int
	var trabalhaAtualmente bool
	var tempoDeEmprego int
	var salario float64

	fmt.Println("Por gentileza, digite a sua idade:")
	fmt.Scanln(&idade)

	fmt.Println("Você trabalha atualmente?")
	fmt.Scanln(&trabalhaAtualmente)

	fmt.Println("Digite o seu tempo de emprego:")
	fmt.Scanln(&tempoDeEmprego)

	fmt.Println("Digite o seu salário:")
	fmt.Scanln(&salario)

	if idade >= 22 && trabalhaAtualmente && tempoDeEmprego > 1 {
		if salario > 500 {
			fmt.Println("Parabens! Seu emprestimo foi aprovado! Sem juros")
		} else {
			fmt.Println("Parabens! Seu emprestimo foi aprovado! Porém, com juros")
		}
	} else {
		fmt.Println("Desculpe, sua solicitacao de emprestimo foi negada :(")
	}
}
