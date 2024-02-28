package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovados := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovados)
	}
}

func main() {
	// não funciona com array
	// aprovados := [4]string{"Maria", "Pedro", "Rafael", "Murilo"}
	aprovados := []string{"Maria", "Pedro", "Rafael", "Murilo"}
	imprimirAprovados(aprovados...)
}
