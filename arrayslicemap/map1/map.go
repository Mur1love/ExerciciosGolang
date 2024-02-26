package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[123654987] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123654987])
	delete(aprovados, 123654987)
	fmt.Println(aprovados[123654987])
}
