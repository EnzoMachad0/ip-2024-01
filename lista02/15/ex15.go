package main

import "fmt"

var comparacoes int
var n1, n2, nA, nB, nInvertido int
var guardaNInvertidos []int

func main() {

	fmt.Printf("ENTRADA\n")
	fmt.Scanln(&comparacoes)
	guardaNInvertidos = make([]int, comparacoes)

	for i := 0; i < comparacoes; i++ {
		fmt.Scanf("%d %d \n", &n1, &n2)

		n1Invertido := inverteNumero(n1)
		n2Invertido := inverteNumero(n2)

		if n1Invertido > n2Invertido {
			guardaNInvertidos[i] = n1Invertido
		} else {
			guardaNInvertidos[i] = n2Invertido
		}

	}
	fmt.Printf("\nSAIDA\n")
	for i := 0; i < comparacoes; i++ {
		fmt.Println(guardaNInvertidos[i])
	}

}

func inverteNumero(n int) int {
	centena := n / 100
	dezena := n % 100 / 10
	unidades := n % 10

	nInvertido := unidades*100 + dezena*10 + centena

	return nInvertido
}
