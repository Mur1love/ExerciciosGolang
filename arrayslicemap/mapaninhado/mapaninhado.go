package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Guga":      1234.5,
			"Guilherme": 12312.6,
		},
		"J": {
			"José": 123145.5,
		},
		"P": {
			"Pedro": 5000.6,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
