package main

import (
	"fmt"
	"os"
)

func estaEmOrdemCrescente(seq []float64) bool {
	n := len(seq)
	for i := 1; i < n; i++ {
		if seq[i-1] >= seq[i] {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("sequencias.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var tamanho int
	var numero float64

	for {
		_, err := fmt.Fscanf(file, "%d\n", &tamanho)
		if err != nil {
			break
		}

		seq := make([]float64, tamanho)
		for i := 0; i < tamanho; i++ {
			fmt.Fscanf(file, "%f", &numero)
			seq[i] = numero
		}

		if estaEmOrdemCrescente(seq) {
			fmt.Println("A sequência está em ordem crescente.")
		} else {
			fmt.Println("A sequência NÃO está em ordem crescente.")
		}
	}
}
