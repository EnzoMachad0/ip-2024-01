package main

import "fmt"

var n int

func main() {
	fmt.Println("ENTRADA: ")
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&arr[i])
	}
	numeroRepetido, _ := repetido(arr)
	_, xRepeticoes := repetido(arr)

	fmt.Println("SAIDA:")

	fmt.Println(numeroRepetido)
	fmt.Println(xRepeticoes)

}

func repetido(arr []int) (numeroRepetido int, xRepeticoes int) {
	cont := 0
	repetidos := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				cont++
			}
		}
		repetidos[arr[i]] = cont
		cont = 0
	}

	var maiorKey int
	var maiorVal int
	for key, val := range repetidos {
		if val > maiorVal {
			maiorKey = key
			maiorVal = val
		} else if key == maiorKey && val < maiorVal {
			val = maiorVal
			key = maiorKey

		}
	}

	return maiorKey, maiorVal
}
