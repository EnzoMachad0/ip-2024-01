package main

import "fmt"

var tamanhoDoVetor, numeroDeTeste, k int

func main() {
	fmt.Println("Digite o tamanho do vetor: ")
	fmt.Scan(&tamanhoDoVetor)

	for tamanhoDoVetor < 1 || tamanhoDoVetor > 100000 {
		fmt.Println("O tamanho do vetor deve ser um valor entre 1 e 100000")
		fmt.Scan(&tamanhoDoVetor)

	}

	fmt.Println("Digite os valores do vetor: ")
	vetor := make([]int, tamanhoDoVetor)
	for i := 0; i < tamanhoDoVetor; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Println("Digite o numero de teste:")
	fmt.Scan(&numeroDeTeste)
	k = numeroDeTeste

	fmt.Println("Saida:")
	fmt.Println(checaNumerosMaiores(vetor, k))

}

func checaNumerosMaiores(vetor []int, k int) int {
	cont := 0
	for i := 0; i < len(vetor); i++ {
		if vetor[i] <= k {
			cont++
		}
	}
	return cont

}
