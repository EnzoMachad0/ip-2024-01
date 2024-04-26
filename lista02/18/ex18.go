package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int) int {
	return a * b / mdc(a, b)
}

func mmcTresNumeros(a, b, c int) int {

	mmcAB := mmc(a, b)

	mmcTotal := mmc(mmcAB, c)
	return mmcTotal
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	mmcTotal := mmcTresNumeros(a, b, c)

	fmt.Printf("%d %d %d :%d\n", a, b, c, mmcTotal)
}
