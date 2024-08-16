package main

import "fmt"

//funções variádicas, podem receber vários parametros

func calculos(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := calculos(1, 3, 5, 6, 7)
	fmt.Println(total)

	escrever("Olá mundo", 1, 2, 4, 5)
}
