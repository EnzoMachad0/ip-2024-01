package main

import (
	"fmt"
)

var matricula int
var hTrabalhadas, vHora, salario float64
var guardaSalario []float64
var guardaMatricula []int

func main() {
	fmt.Println("Digite a matrícula, horas trabalhadas, valor da hora e o salário: ")
	guardaMatricula := make([]int, 0)
	guardaSalario := make([]float64, 0)
	for i := 0; i < 100000; i++ {
		fmt.Scanf("%d %f %f %f \n", &matricula, &hTrabalhadas, &vHora, &salario)

		salario = hTrabalhadas * vHora

		guardaMatricula = append(guardaMatricula, matricula)
		guardaSalario = append(guardaSalario, salario)

		if matricula == 0 {
			break
		}
	}

	fmt.Println("Matrícula e salário: ")

	for i := 0; i < len(guardaSalario)-1; i++ {
		fmt.Println(guardaMatricula[i], guardaSalario[i])
	}

}
