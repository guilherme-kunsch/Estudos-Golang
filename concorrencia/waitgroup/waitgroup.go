package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	//Queremos uma função dependa da outra terminar
	var waitGroup sync.WaitGroup

	//grupo de espera
	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("frase dois")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}