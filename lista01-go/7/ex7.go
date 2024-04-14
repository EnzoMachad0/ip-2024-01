package main

import "fmt"

func main() {
	var tempF, cP, mm, tempC float64

	fmt.Scan(&tempF)
	fmt.Scan(&cP)

	tempC = (5*tempF - 160) / 9

	mm = cP * 25.4

	fmt.Print(" VALOR EM CELSIUS = ", tempC, " A QUANTIDADE DE CHUVA É =", mm)

}
