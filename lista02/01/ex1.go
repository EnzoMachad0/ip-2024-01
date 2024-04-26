package main

import "fmt"

var matricula int
var p1, p2, p3, p4, p5, p6, p7, p8, l1, l2, l3, l4, l5, nt, pres, presenca, notaProva, notaLista, nf float64
var guardaMatricula []int
var guardaNF []float64
var guardaPresenca []float64

func main() {

	guardaMatricula = make([]int, 0)
	guardaNF = make([]float64, 0)
	guardaPresenca = make([]float64, 0)

	fmt.Println("ENTRADA")

	for i := 0; i < 100; i++ {
		fmt.Scanf("%d %f %f %f %f %f %f %f %f %f %f %f %f %f %f %f", &matricula, &p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &l1, &l2, &l3, &l4, &l5, &nt, &pres)

		presenca = pres / 128
		notaProva = (p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8) / 8
		notaLista = (l1 + l2 + l3 + l4 + l5) / 5
		nf = 0.7*notaProva + 0.15*notaLista + 0.15*nt

		guardaPresenca = append(guardaPresenca, presenca)
		guardaMatricula = append(guardaMatricula, matricula)
		guardaNF = append(guardaNF, nf)

		if matricula == -1 {
			break
		}
	}

	fmt.Println("SAÍDA")

	for i := 0; i < len(guardaMatricula)-1; i = i + 2 {
		fmt.Printf("MATRICULA: %d   NOTA FINAL %.2f   SITUAÇÃO FINAL: %s \n", guardaMatricula[i], guardaNF[i], situacaoFinal(guardaNF[i], guardaPresenca[i]))

	}

}

func situacaoFinal(nf float64, presenca float64) string {
	if nf >= 6 && presenca >= 0.75 {
		return "Aprovado"
	} else if nf < 6 && presenca >= 0.75 {
		return "Reprovado por nota"
	} else if nf > 6 && presenca < 0.75 {
		return "Reprovado por frequencia"
	} else {
		return "Reprovado por nota e frequencia"
	}

}
