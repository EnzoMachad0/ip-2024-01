package main

import (
	"fmt"
	"math"
)

func main() {
	var hA, hL, valorDoPagamento float64
	var h float64
	fmt.Scan(&h)

	hA = math.Mod(h, 3) * 5 // mod mostra o resto da divisao h/3

	hL = ((h - hA) / 3) * 10

	valorDoPagamento = hA + hL

	fmt.Print("O VALOR A PAGAR É = ", valorDoPagamento)
}
