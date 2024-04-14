package main

import (
	"fmt"
)

func main() {
	var r, h, aLata, pLata float64
	pi := 3.14159

	fmt.Scan(&r, &h)

	aLata = 2*(pi*r*r) + (2 * pi * r * h)
	pLata = aLata * 100

	fmt.Print("O VALOR DO CUSTO É = ", pLata)

}
