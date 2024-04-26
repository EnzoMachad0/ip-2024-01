package main

import "fmt"

func main() {
	var valorIngresso, valorInicial, valorFinal float64

	// Entrada de dados
	fmt.Println("Digite o valor do ingresso:")
	fmt.Scanln(&valorIngresso)
	fmt.Println("Digite o valor inicial do intervalo:")
	fmt.Scanln(&valorInicial)
	fmt.Println("Digite o valor final do intervalo:")
	fmt.Scanln(&valorFinal)

	// Verificação do intervalo
	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	// Inicialização das variáveis para o resumo
	melhorValorFinal := 0.0
	melhorLucro := 0.0
	numeroIngressos := 0

	// Iteração sobre os valores do intervalo
	for valor := valorInicial; valor <= valorFinal; valor += 1.0 {
		// Cálculo das vendas esperadas
		vendasEsperadas := 120 - (valor-valorIngresso)*20

		lucro := vendasEsperadas*valor - (200.0 + 0.05*vendasEsperadas)

		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorValorFinal = valor
			numeroIngressos = int(vendasEsperadas)
		}

		fmt.Printf("V: %.2f, N: %f, L: %.2f\n", valor, vendasEsperadas, lucro)
	}

	fmt.Printf("Melhor valor final: %.2f\n", melhorValorFinal)
	fmt.Printf("Lucro: %.2f\n", melhorLucro)
	fmt.Printf("Numero de ingressos: %d\n", numeroIngressos)
}
