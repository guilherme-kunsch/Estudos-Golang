package main

import "fmt"

func init() {
	fmt.Println("Executando função init, ela sempre vai executar primeiro")
}

func main() {
	fmt.Println("Função sendo executada!!")
}
