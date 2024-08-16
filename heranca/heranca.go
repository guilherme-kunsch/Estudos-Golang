package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float32
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
	periodo   int
}

func main() {
	fmt.Println("Herança")
	pessoa := Pessoa{
		nome:      "Guilherme",
		sobrenome: "Kunsch",
		idade:     25,
		altura:    1.75,
	}

	estudante := Estudante{
		Pessoa:    pessoa,
		curso:     "ADS",
		faculdade: "UCL",
		periodo:   4,
	}

	fmt.Println(pessoa)
	fmt.Println(estudante)

}
