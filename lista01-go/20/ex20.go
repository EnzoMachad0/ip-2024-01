package main

import "fmt"

func main() {
	var h, m, s, hs, ms, tempoConvertido int

	fmt.Scan(&h, &m, &s)

	hs = h * 3600
	ms = m * 60
	tempoConvertido = hs + ms + s
	fmt.Println("O TEMPO CONVERTIDO EM SEGUNDOS E = ", tempoConvertido)

}
