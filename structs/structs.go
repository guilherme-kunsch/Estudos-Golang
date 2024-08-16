package main

import "fmt"

//structs em Golang é uma coleção de campos

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
	endereco  Endereco
}

type Endereco struct {
	rua    string
	numero int
}

func main() {

	pessoa1 := Pessoa{
		nome:      "Guilherme",
		sobrenome: "Kunsch",
		idade:     25,
		endereco: Endereco{
			rua:    "Realeza",
			numero: 147,
		},
	}

	fmt.Println(pessoa1)

}
