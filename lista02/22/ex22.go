package main

import (
	"fmt"
)

// Função para calcular o máximo divisor comum (MDC)
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Função para converter um número decimal em fração simplificada
func decimalParaFracao(decimal float64) string {
	// Calcular o denominador
	den := 1
	for decimal != float64(int(decimal)) {
		decimal *= 10
		den *= 10
	}
	num := int(decimal)

	mdcNumDen := mdc(num, den)

	num /= mdcNumDen
	den /= mdcNumDen

	return fmt.Sprintf("%d/%d", num, den)
}

func main() {
	var decimal float64
	fmt.Scan(&decimal)

	fração := decimalParaFracao(decimal)
	fmt.Println(fração)
}
