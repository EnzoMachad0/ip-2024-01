package main

import "fmt"

func main() {

	var n1, n2, n3 int
	var x int

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	x = (n1 + n2 + n3) / 3

	if x >= 6 {
		fmt.Print(" Nota = ", x, "APROVADO")
	} else {
		fmt.Print("Nota = ", x, "REPROVADO")
	}

}
