package main

import (
	"fmt"
	"math"
)

var n int

func main() {
	fmt.Println("ENTRADA")

	fmt.Scanln(&n)
	arr := make([]float64, 3*n)

	for i := 0; i < len(arr); i += 3 {
		fmt.Scan(&arr[i], &arr[i+1], &arr[i+2])
	}

	guardaDis := make([]float64, 0)

	for i := 0; i < 3*(n-1); i += 3 {
		d := math.Sqrt(math.Pow((arr[i]-arr[i+3]), 2) + math.Pow((arr[i+1]-arr[i+4]), 2) + math.Pow((arr[i+2]-arr[i+5]), 2))
		guardaDis = append(guardaDis, d)
	}
	fmt.Println("SAÍDA")
	for i := 0; i < len(guardaDis); i++ {
		fmt.Printf("%.2f \n", guardaDis[i])

	}

}
