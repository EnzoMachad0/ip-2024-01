package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var estudantesQueChegaram []int

	for i := 0; i < N; i++ {
		var arrivalTime int
		fmt.Scan(&arrivalTime)
		if arrivalTime <= 0 {
			estudantesQueChegaram = append(estudantesQueChegaram, i+1)
		}
	}

	if len(estudantesQueChegaram) < K {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		for i := len(estudantesQueChegaram) - 1; i >= 0; i-- {
			fmt.Printf("%d ", estudantesQueChegaram[i])
		}
		fmt.Println()
	}
}
