package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar execução!")
	recover()
}

func mediaAluno(n1, n2 float32) bool {

	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6")
}

func main() {
	println(mediaAluno(6, 6))
}
