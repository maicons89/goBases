//Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
//depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
//imposto de um salário.
//Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
//descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
//caso será descontado mais 10%.

package main

import "fmt"

func calcularImposto(salario float64) float64 {
	var imposto float64

	if salario == 50000 {
		imposto = salario * 0.17
	} else {
		imposto = salario * 0.27
	}

	return imposto
}

func main() {
	salarioFuncionario1 := 50000
	salarioFuncionario2 := 150000

	impostoFuncionario1 := calcularImposto(float64(salarioFuncionario1))
	impostoFuncionario2 := calcularImposto(float64(salarioFuncionario2))

	fmt.Printf("Imposto do funcionário 1: R$%.2f\n", impostoFuncionario1)
	fmt.Printf("Imposto do funcionário 2: R$%.2f\n", impostoFuncionario2)

}
