//Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
//Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
//declarar diferentes variáveis.
//Ajude o professor com as seguintes questões:
//1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
//2. Corrigir as incorrectas.
//var 1nome string
//var sobrenome string
//var int idade
//1sobrenome := 6
//var licenca_para_dirigir = true
//var estatura da pessoa int
//quantidadeDeFilhos := 2

package main

import "fmt"

func main() {

	var nome string = "Maicon"
	var idade int = 33
	sobrenome := "Santos"
	var licencaParaDirigir = true
	var estaturaDaPessoa float64 = 1.79
	quantidadeDeFilhos := 2

	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Idade:", idade)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Licença para dirigir:", licencaParaDirigir)
	fmt.Println("Estatura da pessoa:", estaturaDaPessoa)
	fmt.Println("Quantidade de Filhos:", quantidadeDeFilhos)

}
