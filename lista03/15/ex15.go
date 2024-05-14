package main

import (
	"fmt"
	"math"
)

func main() {
	var T, N int
	fmt.Println("ENTRADA: ")
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		fmt.Scan(&N)
		arr := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&arr[i])
		}

		minDistance, comparacoes := minDistancia(arr)
		fmt.Println("SAIDA")
		fmt.Printf("%d %d \n", minDistance, comparacoes)

	}
}

func minDistancia(sequence []int) (int, int) {
	minDistance := math.MaxInt64
	comparacoes := 0

	for i := 0; i < len(sequence); i++ {
		for j := i + 1; j < len(sequence); j++ {
			comparacoes++
			distance := int(math.Abs(float64(sequence[i] - sequence[j])))
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance, comparacoes
}
