package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Murilo":   1235.56,
		"Gulherme": 4123.56,
		"Beatriz":  51231.56,
	}

	funcsESalarios["Rafael"] = 400.45
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
