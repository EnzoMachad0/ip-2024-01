package main

import (
	"fmt"
	"math"
)

func main() {
	var codigo uint64
	var precoCompra, precoVenda float64
	var numVendas, totalVendas, totalCompras int
	var lucroMenor10, lucroEntre10e20, lucroMaior20, maiorLucro, codMaiorVendas uint64

	maiorLucro = 0
	codMaiorVendas = 0

	for {
		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &numVendas)
		if err != nil {
			break
		}

		totalCompras += int(precoCompra * float64(numVendas))
		totalVendas += int(precoVenda * float64(numVendas))

		lucro := ((precoVenda - precoCompra) / precoCompra) * 100

		if lucro < 10 {
			lucroMenor10++
		} else if lucro >= 10 && lucro <= 20 {
			lucroEntre10e20++
		} else {
			lucroMaior20++
		}

		if lucro > float64(maiorLucro) {
			maiorLucro = uint64(math.Round(lucro))
			codMaiorVendas = codigo
		}
	}

	lucroTotal := ((float64(totalVendas) - float64(totalCompras)) / float64(totalCompras)) * 100

	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", lucroEntre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", lucroMaior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codMaiorVendas)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codMaiorVendas)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", float64(totalCompras), float64(totalVendas), lucroTotal)
}
