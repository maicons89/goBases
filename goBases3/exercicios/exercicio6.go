//Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
//detalhes dos dados de cada um deles, conforme o exemplo abaixo:
//Nome: [Nome do aluno]
//Sobrenome: [Sobrenome do aluno]
//RG: [RG do aluno]
//Data de admissão: [Data de admissão do aluno]
//Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
//alunos.
//Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
//Data e que tenha um método de detalhamento

package main

import (
	"fmt"
)

type Aluno struct {
	Nome      string
	Sobrenome string
	RG        string
	Data      string
}

func (a Aluno) Detalhar() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Sobrenome:", a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("Data de admissão:", a.Data)
}

func main() {
	aluno1 := Aluno{
		Nome:      "Maicon",
		Sobrenome: "Santos",
		RG:        "123456789",
		Data:      "12/07/2023",
	}

	aluno2 := Aluno{
		Nome:      "Isa",
		Sobrenome: "Souza",
		RG:        "987654321",
		Data:      "16/03/2023",
	}

	aluno1.Detalhar()
	fmt.Println("-------------------------")
	aluno2.Detalhar()
}
