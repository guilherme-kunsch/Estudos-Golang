package main

import "fmt"

func calculo(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculo(4, 10)
	fmt.Println(soma, subtracao)
}
