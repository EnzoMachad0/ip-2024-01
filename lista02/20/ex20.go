package main

import (
	"fmt"
)

func Perfeito(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

func verificarNumero(n int) {
	sum := 1
	divisores := "1"
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			divisores += fmt.Sprintf(" + %d", i)
			if i != n/i {
				sum += n / i
				divisores += fmt.Sprintf(" + %d", n/i)
			}
		}
	}

	if sum == n {
		fmt.Printf("%d = %s = %d (NUMERO PERFEITO)\n", n, divisores, sum)
	} else {
		fmt.Printf("%d = %s = %d (NUMERO NAO E PERFEITO)\n", n, divisores, sum)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	verificarNumero(n)
}
