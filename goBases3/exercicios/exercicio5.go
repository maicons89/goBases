//Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
//estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
//é que haja muito mais animais para abrigar.
//1. Cães precisam de 10 kg de alimento
//2. Gatos 5 kg
//3. Hamster 250 gramas.
//4. Tarântula 150 gramas.

//Precisamos:
//1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
//com o animal especificado e que retorne uma função com uma mensagem (caso não
//exista o animal)
//2. Uma função para cada animal que calcule a quantidade de alimento com base na
//quantidade necessária do animal digitado.
//Exemplo:

package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

type AnimalFunc func(int) (float64, error)

func Animal(animalType string) (AnimalFunc, error) {
	switch animalType {
	case dog:
		return func(quantity int) (float64, error) {
			foodNeeded := 10.0 * float64(quantity)
			return foodNeeded, nil
		}, nil

	case cat:
		return func(quantity int) (float64, error) {
			foodNeeded := 5.0 * float64(quantity)
			return foodNeeded, nil
		}, nil

	case hamster:
		return func(quantity int) (float64, error) {
			foodNeeded := 0.25 * float64(quantity)
			return foodNeeded, nil
		}, nil

	case tarantula:
		return func(quantity int) (float64, error) {
			foodNeeded := 0.15 * float64(quantity)
			return foodNeeded, nil
		}, nil

	default:
		return nil, errors.New("Animal não encontrado")
	}
}

func main() {
	animalDog, err := Animal(dog)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	animalCat, err := Animal(cat)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Exemplo de uso
	var amount float64
	amountDog, err := animalDog(5)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	amount += amountDog

	amountCat, err := animalCat(8)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	amount += amountCat

	fmt.Printf("Quantidade total de comida necessária: %.2f kg\n", amount)
}
