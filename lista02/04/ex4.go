package main

import (
	"fmt"
)

var n, i, k, s, IS, c, r0, r1 float64

func main() {

	fmt.Println("digite as variaveis n, i, k, s:")
	fmt.Scan(&n, &i, &k, &s)

	Soma()
	Subitrai()
	Multiplica()
	Dividi()

}

func Soma() {

	r0 = n + i
	fmt.Println(" Tabuada de soma:")
	fmt.Printf("%.2f  +  %.2f  =  %.f \n", n, i, r0)

	for c = 1; c < k; c++ {
		r1 = n + i + i*s
		IS = i + i*s
		fmt.Printf("%.2f  +  %.2f  =  %.2f \n", n, IS, r1)

	}

}

func Subitrai() {

	r0 = n - i
	fmt.Println(" Tabuada de subtracao:")

	fmt.Printf("%.2f  -  %.2f  =  %.2f \n", n, i, r0)

	for c = 1; c < k; c++ {
		r1 = n - (i + (i * s))
		IS = i * s
		fmt.Printf("%.2f  -  %.2f  =  %.2f \n", n, IS, r1)

	}

}

func Multiplica() {

	r0 = n * i
	fmt.Println(" Tabuada de multiplicacao:")
	fmt.Printf("%.2f  x  %.2f  =  %.2f \n", n, i, r0)

	for c = 1; c < k; c++ {
		r1 = n * (i + (i * s))
		IS = i + i*s
		fmt.Printf("%.2f  x  %.2f  =  %.2f \n", n, IS, r1)

	}

}
func Dividi() {

	r0 = n / i

	fmt.Println(" Tabuada de divisao:")

	fmt.Printf("%.2f  /  %.2f  =  %.2f \n", n, i, r0)

	for c = 1; c < k; c++ {
		r1 = n / (i + (i * s))
		IS = i * s
		fmt.Printf("%.2f  /  %.2f  =  %.2f \n", n, IS, r1)

	}

}
