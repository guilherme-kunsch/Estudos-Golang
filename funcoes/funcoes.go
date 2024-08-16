package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 15)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Ola man")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(2, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
