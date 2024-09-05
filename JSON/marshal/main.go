package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     int    `json:"idade"`
	Curso     string `json:"curso"`
}

func main() {
	pessoa1 := pessoa{
		Nome:      "Teste",
		Sobrenome: "Testando",
		Idade:     29,
		Curso:     "ADS",
	}

	ret, err := json.Marshal(pessoa1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(ret))

	c2 := map[string]string{
		"nome":  "teste",
		"teste": "nome",
	}

	test, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(test))

}
