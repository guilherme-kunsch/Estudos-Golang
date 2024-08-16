package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiro = referência de memória
	var variavel3 int = 100
	var ponteiro *int //criando um ponteiro

	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

}
