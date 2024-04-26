package main

import (
	"fmt"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func raizQuadrada(n, e float64) {
	r := 1.0

	for abs(n-r*r) > e {
		erro := abs(n - r*r)
		fmt.Printf("r: %.9f, erro: %.9f\n", r, erro)
		r = (r + n/r) / 2
	}

	erro := abs(n - r*r)
	fmt.Printf("r: %.9f, erro: %.9f\n", r, erro)
}

func main() {
	var n, e float64
	fmt.Println("Digite o número cuja raiz quadrada deseja-se obter:")
	fmt.Scanln(&n)
	fmt.Println("Digite o erro considerado (exemplo: 0.00001):")
	fmt.Scanln(&e)

	raizQuadrada(n, e)
}
