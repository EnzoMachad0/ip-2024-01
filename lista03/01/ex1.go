package main

import "fmt"

var numeroDeValoresArr, numeroDeTestes, x int

func main() {
	fmt.Println("digite o numero de valores do array:")

	fmt.Scan(&numeroDeValoresArr)

	arr := make([]int, numeroDeValoresArr)

	fmt.Println("digite os numeros: ")

	for i := 0; i < numeroDeValoresArr; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("digite a quantidade de numeros a serem encontrados:")
	fmt.Scan(&numeroDeTestes)
	fmt.Println("digite os numeros a serem encontrados:")
	arr2 := make([]string, numeroDeTestes)

	for i := 0; i < numeroDeTestes; i++ {

		fmt.Scan(&x)
		if algoritimoDeBusca(arr, x) == true {
			arr2 = append(arr2, "Achei")
		} else {
			arr2 = append(arr2, "Não achei")
		}
	}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

}

func algoritimoDeBusca(arr []int, x int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false

}
