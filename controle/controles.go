package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	}
}
