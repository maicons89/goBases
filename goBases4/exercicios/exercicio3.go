//Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
//que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
//que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
//para as funções:
//- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
//senha
//E devem implementar as seguintes funcionalidades:
//- mudar o nome: me permite mudar o nome e sobrenome
//- mudar a idade: me permite mudar a idade
//- mudar e-mail: me permite mudar o e-mail
//- mudar a senha: me permite mudar a senha

package main

import (
	"fmt"
)

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Senha     string
}

func (u *Usuario) mudarNome(nome, sobrenome string) {
	u.Nome = nome
	u.Sobrenome = sobrenome
}

func (u *Usuario) mudarIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) mudarEmail(email string) {
	u.Email = email
}

func (u *Usuario) mudarSenha(senha string) {
	u.Senha = senha
}

func main() {
	user := Usuario{
		Nome:      "Maicon",
		Sobrenome: "Santos",
		Idade:     36,
		Email:     "maicao@example.com",
		Senha:     "password123",
	}

	fmt.Println("Usuário original:")
	fmt.Println(user)

	// Alterando informações do usuário
	user.mudarNome("Isa", "Dora")
	user.mudarIdade(26)
	user.mudarEmail("dorinha@example.com")
	user.mudarSenha("newpassword456")

	fmt.Println("\nUsuário modificado:")
	fmt.Println(user)
}
