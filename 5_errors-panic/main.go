package main

import (
	"errors"
	"fmt"
	"os"
)

func openFile(fileName string) (*os.File, error) {
	if fileName == "" {
		return nil, errors.New("O nome do arquivo não foi informado")
	}

	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	if fileName == "README.md" {
		panic("Arquivo corrompido!")
	}

	return file, nil
}

func main() {
	file, err := openFile("main.go")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer file.Close()

	fmt.Println("Arquivo encontrado:", file.Name())

	file, err = openFile("")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	file, err = openFile("README.md")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
}
