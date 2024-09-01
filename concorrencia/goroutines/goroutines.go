package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutines são funções ou métodos executados em concorrência
	go escrever("Olá mundo") //goroutine
	escrever("frase dois")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}