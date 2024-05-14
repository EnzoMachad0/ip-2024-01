package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Println("ENTRADA: ")

	fmt.Scan(&N)
	arr := make([]int, 0)

	var anteriorNum, atualNum int
	fmt.Scan(&atualNum)
	arr = append(arr, atualNum)

	for i := 1; i < N; i++ {
		fmt.Scan(&atualNum)
		if atualNum != anteriorNum {
			arr = append(arr, atualNum)
		}
		anteriorNum = atualNum
	}

	fmt.Println("SAIDA: ")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
