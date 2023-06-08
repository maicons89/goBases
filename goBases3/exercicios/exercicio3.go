//Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
//de horas trabalhadas por mês e na categoria do funcionário.
//Se a categoria for C, seu salário é de R$1.000 por hora
//Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
//mensais
//Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
//mensais

//Calcule o salário dos funcionários conforme as informações abaixo:
//- Funcionário de categoria C: 162h mensais
//- Funcionário de categoria B: 176h mensais
//- Funcionário de categoria A: 172h mensais

package main

import "fmt"

const (
	salarioC    = 1000
	salarioB    = 1500
	salarioA    = 3000
	limiteHoras = 160
	bonusB      = 0.2
	bonusA      = 0.5
)

func calcularSalario(categoria string, horasTrabalhadas int) float64 {
	var salarioHora float64

	switch categoria {
	case "C":
		salarioHora = salarioC
	case "B":
		salarioHora = salarioB
		if horasTrabalhadas > limiteHoras {
			horasExtras := horasTrabalhadas - limiteHoras
			bonus := salarioB * bonusB
			salarioHora += bonus * float64(horasExtras)
		}
	case "A":
		salarioHora = salarioA
		if horasTrabalhadas > limiteHoras {
			horasExtras := horasTrabalhadas - limiteHoras
			bonus := salarioA * bonusA
			salarioHora += bonus * float64(horasExtras)
		}
	default:
		fmt.Println("Categoria inválida.")
		return 0
	}

	salario := salarioHora * float64(horasTrabalhadas)
	return salario
}

func main() {
	categoriaC := "C"
	horasTrabalhadasC := 162
	salarioC := calcularSalario(categoriaC, horasTrabalhadasC)
	fmt.Printf("Salário do funcionário C: R$%.2f\n", salarioC)

	categoriaB := "B"
	horasTrabalhadasB := 176
	salarioB := calcularSalario(categoriaB, horasTrabalhadasB)
	fmt.Printf("Salário do funcionário B: R$%.2f\n", salarioB)

	categoriaA := "A"
	horasTrabalhadasA := 172
	salarioA := calcularSalario(categoriaA, horasTrabalhadasA)
	fmt.Printf("Salário do funcionário A: R$%.2f\n", salarioA)
}
