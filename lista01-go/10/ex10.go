package main

import "fmt"

func main() {
	var a, b, c, d, det float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	det = a*d - b*c

	fmt.Print("O VALOR DO DETERMINANTE E = ", det)
}
