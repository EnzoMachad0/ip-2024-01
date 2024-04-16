package main

import "fmt"

func main() {
	var n int
	var soma float64
	fmt.Println("Digite o número de termos da sequência: ")
	fmt.Scan(&n)

	if n < 1 {

		fmt.Println("Número de termos inválido")

	} else {

		for i := 1; i <= n; i++ {

			soma += 1 / float64(i)

		}

		fmt.Println("A soma dos termos da sequência é: ", soma)

	}
}
