package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{
		Name:  "Eleson",
		Age:   27,
		Email: "eleson@mail.com",
	}

	personJSON, err := Serialize(person)

	if err != nil {
		log.Fatalf("Erro ao serializar: %s", err)
	}

	personStruct, err := Deserialize[Person](personJSON)

	if err != nil {
		log.Fatalf("Erro ao deserializar: %s", err)
	}

	fmt.Println("JSON Serializado:", personJSON)
	fmt.Println("JSON Deserializado:", *personStruct)
}

func Serialize(data any) (string, error) {
	dataAsByte, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(dataAsByte), nil
}

func Deserialize[T any](data string) (*T, error) {
	var dataDeserialized T

	err := json.Unmarshal([]byte(data), &dataDeserialized)

	if err != nil {
		return nil, err
	}

	return &dataDeserialized, nil
}
