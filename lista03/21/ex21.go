package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var pares, impares []int

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
	}

	for i := 0; i < len(pares)-1; i++ {
		for j := i + 1; j < len(pares); j++ {
			if pares[i] > pares[j] {
				pares[i], pares[j] = pares[j], pares[i]
			}
		}
	}

	for i := 0; i < len(impares)-1; i++ {
		for j := i + 1; j < len(impares); j++ {
			if impares[i] < impares[j] {
				impares[i], impares[j] = impares[j], impares[i]
			}
		}
	}

	for _, num := range pares {
		fmt.Println(num)
	}
	for _, num := range impares {
		fmt.Println(num)
	}
}
