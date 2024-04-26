package main

import "fmt"

var numeros []int
var par []int
var impar []int
var mp, mi, sp, si float64
var n int

func main() {
	fmt.Println("Digite uma sequencia de números inteiros: ")

	numeros = make([]int, 0)

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			numeros = append(numeros, n)
		}
	}

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			par = append(par, numeros[i])
		} else {
			impar = append(impar, numeros[i])
		}

	}
	for i := 0; i < len(par); i++ {
		sp += float64(par[i])
	}

	mp = sp / float64(len(par))

	for i := 0; i < len(impar); i++ {
		si += float64(impar[i])
	}

	mi = si / float64(len(impar))

	fmt.Println("Média par: ", mp)
	fmt.Println("Média ímpar: ", mi)

}
