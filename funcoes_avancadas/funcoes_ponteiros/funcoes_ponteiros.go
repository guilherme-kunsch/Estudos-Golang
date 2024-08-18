package main

import "fmt"

func inverteSinal(n1 int) int {
	return n1 * 1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20

	numeroInvertido := inverteSinal(numero)

	fmt.Println(numeroInvertido)

	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
