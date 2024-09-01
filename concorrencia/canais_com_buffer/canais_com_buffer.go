package main

import (
	"fmt"
)

func main() {
	//canal com buffer de 2
	//canal com buffer só vai bloquear quando atingir a capacidade dele
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}