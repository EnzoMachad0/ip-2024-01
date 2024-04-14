package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NÚMERO É ÍMPAR")
	} else {
		fmt.Println(x)
		for i := 1; i < y; i++ {
			fmt.Println(x + 2)
			x = x + 2
		}
	}

}
