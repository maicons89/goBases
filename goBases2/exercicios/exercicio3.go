//Faça uma aplicação que contenha uma variável com o número do mês.
//1. Dependendo do número, imprima o mês correspondente em texto.
//2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
//escolheria e porquê?
//Ex: 7 de julho.

package main

import (
	"fmt"
)

func main() {

	var mes int

	fmt.Println("Digite o número do mês:")
	fmt.Scanln(&mes)

	switch mes {
	case 0:
		fmt.Println("Janeiro")
	case 1:
		fmt.Println("Fevereiro")
	case 2:
		fmt.Println("Março")
	case 3:
		fmt.Println("Abril")
	case 4:
		fmt.Println("Maio")
	case 5:
		fmt.Println("Junho")
	case 6:
		fmt.Println("Julho")
	case 7:
		fmt.Println("Agosto")
	case 8:
		fmt.Println("Setembro")
	case 9:
		fmt.Println("Outubro")
	case 10:
		fmt.Println("Novembro")
	case 11:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Desconhecido")
	}

}
