package main

import "math"

// inicializando com letra maiuscula é PUBLICO (visivel dentro e fora do pacote)
// iniciando com letra minuscula é PRIVADO (Visivel apenas dentro do pacote)

// Por exemplo...
// Ponto - gerará algo publico
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma cordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distancia linear entre os dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
