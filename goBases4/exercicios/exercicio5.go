//Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
//Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
//dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
//eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

//Precisamos de 3 estruturas:
//- Produtos: nome, preço, quantidade.
//- Serviços: nome, preço, minutos trabalhados.
//- Manutenção: nome, preço.
//Precisamos de 3 funções:
//- Somar Produtos: recebe um array de produto e devolve o preço total (preço *
//quantidade).
//- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
//hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
//tivesse trabalhado meia hora).
//- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

//Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
//tela (somando o total dos 3).

package main

import (
	"fmt"
	"sync"
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Servico struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados int
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func SomarProdutos(produtos []*Produto, wg *sync.WaitGroup, totalCh chan<- float64) {
	defer wg.Done()

	total := 0.0
	for _, p := range produtos {
		total += p.Preco * float64(p.Quantidade)
	}

	totalCh <- total
}

func SomarServicos(servicos []*Servico, wg *sync.WaitGroup, totalCh chan<- float64) {
	defer wg.Done()

	total := 0.0
	for _, s := range servicos {
		minutos := float64(s.MinutosTrabalhados)
		if minutos < 30 {
			minutos = 30
		}

		total += s.Preco * (minutos / 60)
	}

	totalCh <- total
}

func SomarManutencao(manutencoes []*Manutencao, wg *sync.WaitGroup, totalCh chan<- float64) {
	defer wg.Done()

	total := 0.0
	for _, m := range manutencoes {
		total += m.Preco
	}

	totalCh <- total
}

func main() {
	produtos := []*Produto{
		{Nome: "Produto 1", Preco: 10.0, Quantidade: 2},
		{Nome: "Produto 2", Preco: 5.0, Quantidade: 3},
	}

	servicos := []*Servico{
		{Nome: "Serviço 1", Preco: 20.0, MinutosTrabalhados: 45},
		{Nome: "Serviço 2", Preco: 15.0, MinutosTrabalhados: 20},
	}

	manutencoes := []*Manutencao{
		{Nome: "Manutenção 1", Preco: 50.0},
		{Nome: "Manutenção 2", Preco: 30.0},
	}

	var wg sync.WaitGroup
	wg.Add(3)

	totalCh := make(chan float64)

	go SomarProdutos(produtos, &wg, totalCh)
	go SomarServicos(servicos, &wg, totalCh)
	go SomarManutencao(manutencoes, &wg, totalCh)

	go func() {
		wg.Wait()
		close(totalCh)
	}()

	total := 0.0
	for val := range totalCh {
		total += val
	}

	fmt.Printf("Preço total: %.2f\n", total)
}
