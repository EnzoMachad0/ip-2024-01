package main

import "fmt"

func main() {

	var nota float64

	fmt.Scan(&nota)

	if nota >= 9 && nota <= 10 {
		fmt.Println("NOTA = ", nota, "CONCEITO = A")
	} else if nota >= 7.5 && nota < 9 {
		fmt.Println("NOTA = ", nota, "CONCEITO = B")
	} else if nota >= 6 && nota < 7.5 {
		fmt.Println("NOTA = ", nota, "CONCEITO = C")
	} else if nota >= 0 && nota < 6 {
		fmt.Println("NOTA = ", nota, "CONCEITO = D")
	} else {
		fmt.Println("NOTA INVÁLIDA, DIGITE UMA NOTA ENTRE 0 E 10.")
	}

}
