package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, tempo float64
	tA := 1.03
	tB := 1.015
	fmt.Println("Digite a pupulação do país A : ")
	fmt.Scan(&a)
	fmt.Println("Digite a pupulação do país B : ")
	fmt.Scan(&b)

	for {
		tempo++
		x := math.Pow(tA, tempo)
		y := math.Pow(tB, tempo)
		c := a * x
		d := b * y

		if c > d {
			fmt.Println("O tempo que a população A ultrapassa a população B é de : ", tempo, "anos")
			break
		}

	}

}

//for para checar o tempo que a população A ultrapassa a população B
