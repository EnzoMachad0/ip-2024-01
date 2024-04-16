package main

import "fmt"

func main() {

	var n, tf int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Print("digite a temperatura em fahrenheit: ")
		fmt.Scan(&tf)
		n = (tf - 32) * 5 / 9
		fmt.Println("A temperatura em celsius é: ", n)
	}

}
