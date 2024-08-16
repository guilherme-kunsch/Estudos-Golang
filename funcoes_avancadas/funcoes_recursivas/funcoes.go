package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Salario   float32
}

func main() {
	var p Pessoa
	fmt.Println("Qual seu nome: ")
	fmt.Scanln(&p.Nome)
	fmt.Println("Qual seu sobrenome: ")
	fmt.Scanln(&p.Sobrenome)
	fmt.Println("Qual sua idade: ")
	fmt.Scanln(&p.Idade)
	fmt.Println("Qual seu salário: ")
	fmt.Scanln(&p.Salario)

	fmt.Println(p)
}
