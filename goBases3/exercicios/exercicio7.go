//Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
//produtos e devolver o valor do preço total.
//As empresas têm 3 tipos de produtos:
//- Pequeno, Médio e Grande.
//Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

//Lista de custos adicionais:
//- Pequeno: O custo do produto (sem custo adicional)
//- Médio: O custo do produto + 3% pela disponibilidade no estoque
//- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
//adicional pelo envio de $2500.

//2

//Requisitos:
//- Criar uma estrutura “loja” que guarde uma lista de produtos.
//- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
//- Criar uma interface “Produto” que possua o método “CalcularCusto”
//- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
//- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
//e preço, e devolva um Produto.
//- Será necessário uma função “novaLoja” que retorne um Ecommerce.
//- Interface Produto:
//- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
//custo adicional segundo o tipo de produto.

//- Interface Ecommerce:
//- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
//base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
//tenha)
//- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
//e adicioná-lo a lista da loja

package main

import "fmt"

type ProdutoTipo string

const (
	Pequeno ProdutoTipo = "Pequeno"
	Medio   ProdutoTipo = "Médio"
	Grande  ProdutoTipo = "Grande"
)

type Produto struct {
	Tipo  ProdutoTipo
	Nome  string
	Preco float64
}

func (p Produto) CalcularCusto() float64 {
	custo := p.Preco

	switch p.Tipo {
	case Medio:
		custo += p.Preco * 0.03
	case Grande:
		custo += p.Preco * 0.06
		custo += 2500.0
	}

	return custo
}

type Ecommerce interface {
	Total() float64
	Adicionar(produto Produto)
}

type Loja struct {
	Produtos []Produto
}

func (l *Loja) Total() float64 {
	total := 0.0

	for _, produto := range l.Produtos {
		total += produto.CalcularCusto()
	}

	return total
}

func (l *Loja) Adicionar(produto Produto) {
	l.Produtos = append(l.Produtos, produto)
}

func novoProduto(tipo ProdutoTipo, nome string, preco float64) Produto {
	return Produto{
		Tipo:  tipo,
		Nome:  nome,
		Preco: preco,
	}
}

func novaLoja() Ecommerce {
	return &Loja{
		Produtos: []Produto{},
	}
}

func main() {
	loja := novaLoja()

	produto1 := novoProduto(Pequeno, "Produto Pequeno 1", 100.0)
	loja.Adicionar(produto1)

	produto2 := novoProduto(Medio, "Produto Médio 1", 200.0)
	loja.Adicionar(produto2)

	produto3 := novoProduto(Grande, "Produto Grande 1", 300.0)
	loja.Adicionar(produto3)

	total := loja.Total()
	fmt.Printf("Preço Total: R$%.2f\n", total)
}
