//Uma empresa que vende produtos de limpeza precisa de:
//1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
//de produtos comprados, separados por ponto e vírgula (csv).
//2. Deve possuir o ID do produto, preço e a quantidade.
//3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Dados dos produtos
	produtos := []struct {
		ID         int
		Preco      float64
		Quantidade int
	}{
		{ID: 1, Preco: 10.99, Quantidade: 5},
		{ID: 2, Preco: 5.99, Quantidade: 10},
		{ID: 3, Preco: 7.5, Quantidade: 3},
	}

	// Abrir arquivo para escrita
	arquivo, err := os.Create("produtos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	// Escrever os produtos no arquivo
	for _, produto := range produtos {
		linha := fmt.Sprintf(" %d; %f; %d \n", produto.ID, produto.Preco, produto.Quantidade)
		_, err := arquivo.WriteString(linha)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Arquivo 'produtos.csv' criado com sucesso!")
}
