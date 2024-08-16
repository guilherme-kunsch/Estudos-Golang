package main

import "fmt"

func main() {
	var teste string = "guilherme"
	variavel := 3

	var (
		variavel3 string = "teste"
		variavel4 string = "123"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel)
	fmt.Println(teste)

	variavel3, variavel4 = variavel4, variavel3

}
