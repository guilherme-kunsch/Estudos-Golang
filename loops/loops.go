package main

import "fmt"

func main() {
	nomes := [3]string{"Gui", "Teste", "Gaby"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "guil",
		"sobrenome": "teste",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
