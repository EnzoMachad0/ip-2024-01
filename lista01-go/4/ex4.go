package main

import "fmt"

func main() {
	var sm, kwg, ckw, cc, cd float64

	fmt.Scan(&sm)
	fmt.Scan(&kwg)

	ckw = 0.007 * sm
	cc = ckw * kwg
	cd = 0.9 * cc

	fmt.Println("Custo por kw: R$", ckw)
	fmt.Println("Custo consumo: R$", cc)
	fmt.Println("Custo com desconto: R$", cd)

}
