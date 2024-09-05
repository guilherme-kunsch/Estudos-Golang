package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     string `json:"idade"`
	Curso     string `json:"curso"`
}

func main() {
	pessoaUm := `{"nome":"Teste","sobrenome":"Testando","idade":"29","curso":"ADS"}`

	var p pessoa

	if erro := json.Unmarshal([]byte(pessoaUm), &p); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(p)

	pessoaDois := `{"nome":"Testando","sobrenome":"Teste","idade":"30","curso":"ADS"}`

	p2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(pessoaDois), &p2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(p2)
}
