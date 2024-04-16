package main

import "fmt"

func main() {
	var a1, r, n, soma int

	fmt.Print("Digite o primeiro termo da PA: ")
	fmt.Scan(&a1)
	fmt.Print("Digite a razão da PA: ")
	fmt.Scan(&r)
	fmt.Print("Digite o número de termos da PA: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		soma += a1 + i*r
	}

	fmt.Println("A soma dos termos da PA é: ", soma)
}
