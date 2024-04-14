package main

import "fmt"

func main() {
	var sf, sfr float64

	fmt.Scan(&sf)

	if sf <= 300 {
		sfr = sf * 1.5
	} else {
		sfr = sf * 1.3
	}

	fmt.Println("SÁLARIO COM REAJUSTE = R$", sfr)

}
