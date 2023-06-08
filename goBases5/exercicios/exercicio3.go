//Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
//erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
//tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
//o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

package main

import (
	"fmt"
)

func main() {
	salario := 12000

	if salario < 15000 {
		err := fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
		panic(err)
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
