package main

import (
	"fmt"
	"strings"
)

func main() {

	palavra := "Maicon"

	numLetras := 0
	for range palavra {
		numLetras++
	}
	fmt.Printf("A palavra '%s' contém %d letras.\n", palavra, numLetras)

	fmt.Println("Imprimi o número de letras:")
	letras := strings.Split(palavra, "")
	for i := 0; i < len(letras); i++ {
		letra := letras[i]
		fmt.Println(letra)
	}

}
