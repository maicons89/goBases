//Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
//produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
//mesmo endereço de memória no main do programa e nas funções.

//struturas necessárias:
//- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
//- Produto: Nome, preço, quantidade.
//Algumas funções necessárias:
//- Novo produto: recebe nome e preço, e retorna um produto.
//- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
//o produto ao usuário.
//- Deletar produtos: recebe um usuário, apaga os produtos do usuário.

package main

import (
	"fmt"
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []*Produto
}

func NovoProduto(nome string, preco float64, quantidade int) *Produto {
	return &Produto{
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}
}

func AdicionarProduto(usuario *Usuario, produto *Produto, quantidade int) {
	produtoExistente := false
	for _, p := range usuario.Produtos {
		if p == produto {
			p.Quantidade += quantidade
			produtoExistente = true
			break
		}
	}

	if !produtoExistente {
		produto.Quantidade = quantidade
		usuario.Produtos = append(usuario.Produtos, produto)
	}
}

func DeletarProdutos(usuario *Usuario) {
	usuario.Produtos = nil
}

func main() {
	usuario := &Usuario{
		Nome:      "Maicon",
		Sobrenome: "Santos",
		Email:     "maicon@example.com",
		Produtos:  []*Produto{},
	}

	produto1 := NovoProduto("Camiseta", 99.99, 2)
	produto2 := NovoProduto("Calça", 159.99, 1)

	AdicionarProduto(usuario, produto1, 2)
	AdicionarProduto(usuario, produto2, 1)

	fmt.Println("Produtos do usuário:")
	for _, p := range usuario.Produtos {
		fmt.Printf("Nome: %s, Preço: %.2f, Quantidade: %d\n", p.Nome, p.Preco, p.Quantidade)
	}

	DeletarProdutos(usuario)

	fmt.Println("\nProdutos do usuário após deletar:")
	for _, p := range usuario.Produtos {
		fmt.Printf("Nome: %s, Preço: %.2f, Quantidade: %d\n", p.Nome, p.Preco, p.Quantidade)
	}
}
