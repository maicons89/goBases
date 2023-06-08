//Um escritório de contabilidade precisa acessar os dados de seus funcionários para
//poder realizar diferentes liquidações. Para isso, eles têm todos os detalhes
//necessários em um arquivo .txt.
//1. É necessário desenvolver a funcionalidade para poder ler o arquivo .txt
//indicado pelo cliente, porém, eles não passaram o arquivo para ser lido pelo
//nosso programa.
//2. Desenvolva o código necessário para ler os dados do arquivo chamado
//“customers.txt” (lembre-se do que você viu sobre o pacote “os”).
//Como não temos o arquivo necessário, será obtido um erro e, neste caso, o programa
//terá que exibir um panic ao tentar ler um arquivo que não existe, exibindo a
//mensagem “o arquivo indicado não foi encontrado ou está danificado ”.

//Requisitos gerais:
//● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
//● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
//um valor de erro (por exemplo, aqueles que tentam ler arquivos).
//Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
//que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
//o caso de erro retornado).

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// CustomError é um erro personalizado que pode ser utilizado para indicar que o arquivo não foi encontrado ou está danificado
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func main() {
	filePath := "customers.txt"

	// Verifica se o arquivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := &CustomError{"O arquivo indicado não foi encontrado ou está danificado"}
		panic(err)
	}

	// Lê o conteúdo do arquivo
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Exibe o conteúdo do arquivo
	fmt.Println(string(content))
}
