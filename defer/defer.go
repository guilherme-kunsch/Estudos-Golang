package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func mediaAluno(n1, n2 float32) bool {
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	fmt.Println("---Defer---")
	defer funcao1()
	funcao2()

}
