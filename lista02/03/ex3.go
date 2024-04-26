package main

import ("fmt"


)

var n, f int

func main() {

	fmt.Println("Digite o numero inteiro: ")
	fmt.Scan(&n)

	fmt.Printf("%d! = %d", n, recursaoFat(n))

}
//função recursiva: chama a si mesma 
func recursaoFat(n int) int {

	if n == 1 {// caso base, retira a possibilidade do loop infinito 
		return 1
	} else {
		return n * recursaoFat(n-1)// chama novamente a func
	}

}


