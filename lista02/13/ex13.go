package main

import "fmt"

func main() {

	var n, total_graos int

	fmt.Println("Digite o número de grãos a ser colocado no primeiro quadrado:")
	fmt.Scanln(&n)

	total_graos = 0

	for i := 1; i <= 64; i++ {

		if i%2 == 0 {
			total_graos += n
		} else {
			total_graos += n
		}

		n *= 2
	}

	fmt.Println("Quantidade total de grãos:", total_graos)
}
