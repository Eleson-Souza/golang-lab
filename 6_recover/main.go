package main

import (
	"fmt"
	"time"
)

func calcAgeByYear(birthYear int) *int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro crítico capturado e recuperado:", r)
			fmt.Println("Seguindo programa...")
		}
	}()

	age := time.Now().Year() - birthYear

	if age < 0 || age > 130 {
		panic("O ano informado é inválido!")
	}

	// mesmo com recovery, essa linha não é executada, em caso de panic.
	return &age
}

func main() {
	fmt.Println("Início do programa")

	userAge := calcAgeByYear(2025)
	if userAge != nil {
		fmt.Printf("\n%d anos\n", *userAge)
	}

	fmt.Println("\nFim do programa")
}
