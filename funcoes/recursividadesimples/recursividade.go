package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	// resultado2 := fatorial(-2) // uint só usa numeros positivos
	// fmt.Println(resultado2)

	//Uma solução melhor seria...
}
