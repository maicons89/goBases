//O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
//cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
//- Arquivo
//- Nome e sobrenome
//- RG
//- Número de telefone
//- Endereço

//● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
//da cobrança das despesas restantes. Desenvolva e implemente uma função para
//gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
//Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
//interrompa a execução e aborte

//● Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
//fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
//código, implemente a função para ler um arquivo chamado “customers.txt” (como no
//exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
//retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
//agora. Esse erro deve:

//1. Gerar um panic;
//2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
//está danificado”, e continuar com a execução do programa normalmente.

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
	"os"
)

// CustomError é um erro personalizado que pode ser utilizado para indicar que o arquivo não foi encontrado ou está danificado
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

// Tarefa 1: Função para gerar um ID de arquivo
func generateFileID() string {
	// Implemente a lógica para gerar o ID do arquivo (exemplo: usando uma biblioteca para gerar um UUID)

	// Se ocorrer algum erro na geração do ID, retorne nil
	return "12345" // Exemplo de ID gerado
}

// Tarefa 2: Função para ler o arquivo customers.txt
func readCustomersFile() error {
	filePath := "customers.txt"

	// Verifica se o arquivo existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := &CustomError{"O arquivo indicado não foi encontrado ou está danificado"}
		panic(err)
	}

	// Lê o conteúdo do arquivo
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Exibe o conteúdo do arquivo
	fmt.Println(string(content))

	return nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recuperado:", r)
		}
	}()

	// Tarefa 1: Geração do ID do arquivo
	fileID := generateFileID()
	if fileID == "" {
		panic("Falha ao gerar o ID do arquivo")
	}

	// Tarefa 2: Leitura do arquivo customers.txt
	err := readCustomersFile()
	if err != nil {
		panic(err)
	}

	// Restante do código para cadastrar o cliente
	fmt.Println("Cadastro do cliente realizado com sucesso!")
}
