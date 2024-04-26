package main

import (
	"fmt"
)

func maiorSubsequenciaIgual(n int, sequencia []int) int {
	maiorSubsequencia := 0
	contador := 1 // Contador para a subsequência atual

	// Iterar pela sequência para encontrar a maior subsequência
	for i := 1; i < n; i++ {
		if sequencia[i] == sequencia[i-1] {
			contador++
		} else {
			if contador > maiorSubsequencia {
				maiorSubsequencia = contador
			}
			contador = 1
		}
	}

	// Atualizar a maior subsequência, caso a última subsequência seja a maior
	if contador > maiorSubsequencia {
		maiorSubsequencia = contador
	}

	return maiorSubsequencia
}

func main() {
	var n int
	fmt.Println("Digite o número de elementos da sequência:")
	fmt.Scanln(&n)

	sequencia := make([]int, n)
	fmt.Println("Digite os elementos da sequência separados por espaço:")
	for i := 0; i < n; i++ {
		fmt.Scan(&sequencia[i])
	}

	tamanhoMaiorSubsequencia := maiorSubsequenciaIgual(n, sequencia)

	fmt.Println("A maior subsequência de elementos iguais possui", tamanhoMaiorSubsequencia, "elementos.")
}
