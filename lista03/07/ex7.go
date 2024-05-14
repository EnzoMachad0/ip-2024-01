package main

import "fmt"

func main() {
	var N int

	for {
		fmt.Scan(&N)
		if N == 0 {
			break
		}

		vetor := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&vetor[i])
		}

		M := 0
		for _, num := range vetor {
			if num > M {
				M = num
			}
		}

		freq := make([]int, M+1)
		for _, num := range vetor {
			freq[num]++
		}

		for i := 0; i <= M; i++ {
			fmt.Printf("(%d) %d\n", i, freq[i])
		}
		fmt.Println()
	}
}
