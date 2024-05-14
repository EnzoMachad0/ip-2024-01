package main

import "fmt"

var n int

func main() {

	guardaIndice := make([]int, 0)
	guardaTheBiggest := make([]int, 0)

	for {
		fmt.Scan(&n)
		if n != 0 {
			vetor := make([]int, n)

			for i := 0; i < len(vetor); i++ {
				fmt.Scan(&vetor[i])
			}

			maior, _ := theBiggest(vetor)
			guardaTheBiggest = append(guardaTheBiggest, maior)
			_, indice := theBiggest(vetor)
			guardaIndice = append(guardaIndice, indice)

		} else {
			break
		}

	}
	for i := 0; i < len(guardaTheBiggest); i++ {
		fmt.Printf("%d %d ", guardaTheBiggest[i], guardaIndice[i])

	}
}

func theBiggest(vetor []int) (maior int, indice int) {
	if len(vetor) == 0 {
		return 0, -1
	}

	maior = vetor[0]
	for i := 1; i < len(vetor); i++ {
		if vetor[i] > maior {
			maior = vetor[i]
			indice = i
		}
	}

	return maior, indice

}
