//1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do tipo “int”.
//2. Crie um erro personalizado com uma struct que implemente “Error()” com a
//mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
//disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
//console a mensagem “Necessário pagamento de imposto”.

package main

import (
	"fmt"
)

type SalarioError struct {
	message string
}

func (e *SalarioError) Error() string {
	return e.message
}

func main() {
	salario := 12000

	if salario < 15000 {
		err := &SalarioError{"erro: O salário digitado não está dentro do valor mínimo"}
		panic(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
