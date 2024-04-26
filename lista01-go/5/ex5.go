package main

import "fmt"

func main() {
	var contaDoCliente, cam3, valorDaConta float64
	var tipoDoConsumidor string
	fmt.Scan(&contaDoCliente)
	fmt.Scan(&cam3)
	fmt.Scan(&tipoDoConsumidor)

	fmt.Println("Digite o número da conta:")
	fmt.Println("Digite o consumo:")
	fmt.Println("Digite o tipo do consumidor:")

	switch tipoDoConsumidor {
	case "r":
		valorDaConta = 5 + 0.05*cam3
	case "d":
		if cam3 <= 800 {
			valorDaConta = 500
		} else {
			valorDaConta = 500 + 0.25*(cam3-80)
		}
	case "i":
		if cam3 <= 100 {
			valorDaConta = 800
		} else {
			valorDaConta = 800 + 0.04*(cam3-100)
		}
	}

	fmt.Print("Conta = ", contaDoCliente, " Valor da conta = ", valorDaConta)

}
