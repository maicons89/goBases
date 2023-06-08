//Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
//De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

//var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

//Por outro lado, você também precisa:
//- Saiba quantos de seus funcionários têm mais de 21 anos.
//- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
//- Excluir Pedro do mapa.

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("A idade de Benjamin é %d\n", employees["Benjamin"])

	quant_funcionario_22 := 0
	for _, idade := range employees {
		if idade > 21 {
			quant_funcionario_22++
		}
	}

	fmt.Println("Numero de funcionário com mais de 21 = ", quant_funcionario_22)

	employees["Frederico"] = 25

	fmt.Println("Adicionando um novo funcionário chamado: ", employees)

	delete(employees, "Pedro")
	fmt.Println("Removendo o funcionário Pedro: ", employees)
}
