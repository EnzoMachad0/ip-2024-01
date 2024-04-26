package main

import "fmt"

var nDeTimes int
var final = 1

func main() {

	fmt.Println("Digite o número de times: ")
	fmt.Scanln(&nDeTimes)
	if nDeTimes <= 1 {
		fmt.Printf("Campeonato inválido.")
	} else {
		combinatoriaSimples(nDeTimes)

		// i é igual a um pois seu valor sera passado como time
		for i := 1; i <= combinatoriaSimples(nDeTimes); i++ {
			for j := i + 1; j <= nDeTimes; j++ {
				fmt.Printf("Final %d: Time %d x Time %d\n", final, i, j)
				final++
			}
		}
	}

}

func fat(nDetimes int) int {
	if nDetimes == 1 {
		return 1
	} else {
		return nDetimes * fat(nDetimes-1)
	}
}

func combinatoriaSimples(nDeTimes int) int {
	return fat(nDeTimes) / (fat(2) * fat(nDeTimes-2))
}
