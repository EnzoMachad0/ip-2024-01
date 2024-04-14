package main

import "fmt"

func main() {
	var n1, n2, n3, nl, ne int

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	nl = n1 * n2 * n3
	ne = nl * nl

	if n1 == 0 && n2 == 0 {
		nl = n3
		ne = nl * nl
		fmt.Println(nl, ne)
	} else if n1 == 0 && n2 != 0 {
		nl = n2 * n3
		ne = nl * nl
		fmt.Println(nl, ne)
	} else if n1 != 0 && n2 == 0 {
		nl = n1 * n3
		ne = nl * nl
		fmt.Println(nl, ne)
	} else if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("DÍGITO INVÁLIDO")
	} else {
		fmt.Println(nl, "e", ne)
	}

}
