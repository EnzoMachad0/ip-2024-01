package main

import "fmt"

func gerarSomaImpares(n int) []int {
	soma := 0
	somas := make([]int, n)

	for i := 1; i <= n; i++ {
		soma += 2*i - 1
		somas[i-1] = soma
	}

	return somas
}

func main() {
	var m int
	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		somas := gerarSomaImpares(i)
		fmt.Printf("%d*%d*%d =", i, i, i)
		for j, soma := range somas {
			if j > 0 {
				fmt.Print("+")
			}
			fmt.Printf("%d", soma)
		}
	}
}
