package main

//canal de comunicação, pode enviar e receber dados, podemos sincronizar goroutines

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	//aqui vai acontecer deadlock
	for {
		//esperando receber um valor
		//passamos uma segunda variável, para ver se o canal está aberto
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

//usando o canal, esse FOR irá executar apenas uma vez
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++{
		//estou mandando um valor para canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//depois que terminar o loop iremos fechar o canal
	close(canal)
}