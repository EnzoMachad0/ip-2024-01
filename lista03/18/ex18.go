package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var cpf [11]int
		for j := 0; j < 11; j++ {
			fmt.Scan(&cpf[j])
		}

		if validarCPF(cpf) {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}

func validarCPF(cpf [11]int) bool {
	if len(cpf) != 11 {
		return false
	}

	var soma1, soma2 int

	for i := 0; i < 9; i++ {
		soma1 += cpf[i] * (10 - i)
		soma2 += cpf[i] * (11 - i)
	}

	soma1 %= 11
	if soma1 < 2 {
		soma1 = 0
	} else {
		soma1 = 11 - soma1
	}

	soma2 += soma1 * 2
	soma2 %= 11
	if soma2 < 2 {
		soma2 = 0
	} else {
		soma2 = 11 - soma2
	}

	return soma1 == cpf[9] && soma2 == cpf[10]
}
