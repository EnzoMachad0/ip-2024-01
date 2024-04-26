package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	sliceDeNúmeros := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sliceDeNúmeros[i])
	}

	comprimentoAtual := 1
	comprimentoMaximo := 1

	for i := 1; i < n; i++ {
		if sliceDeNúmeros[i] < sliceDeNúmeros[i+1] {
			comprimentoAtual++
		} else {

			if comprimentoAtual > comprimentoMaximo {
				comprimentoMaximo = comprimentoAtual
			}
			comprimentoAtual = 1
		}
	}

	if comprimentoAtual > comprimentoMaximo {
		comprimentoMaximo = comprimentoAtual
	}

	fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", comprimentoMaximo-1)
}
