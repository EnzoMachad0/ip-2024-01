package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Print("O numero ", n, " é divisivel por 3 e por 5")
	} else {
		fmt.Print("O numero ", n, " nao é divisivel por 3 e por 5")
	}
}
