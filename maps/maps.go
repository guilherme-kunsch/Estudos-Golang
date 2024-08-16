package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Guilherme",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "ads",
			"campus": "ulc",
		},
	}

	fmt.Println(usuario2)
}
