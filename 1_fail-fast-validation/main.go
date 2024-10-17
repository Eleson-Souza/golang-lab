package main

import (
	"errors"
	"fmt"
	"regexp"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func ValidateUser(user User) error {
	if len(user.Name) < 3 {
		return errors.New("o nome deve ter pelo menos 3 caracteres")
	}

	if matched, _ := regexp.MatchString(`\S+@\S+\.\S+`, user.Email); !matched {
		return errors.New("formato de email inválido")
	}

	if user.Age < 18 {
		return errors.New("o usuário deve ter 18 anos ou mais")
	}

	return nil
}

func main() {
	user := User{
		Name:  "Eleson",
		Email: "eleson@domain.com",
		Age:   27,
	}

	if err := ValidateUser(user); err != nil {
		fmt.Println("Erro de validação:", err)
	} else {
		fmt.Println("Usuário válido!")
	}
}
