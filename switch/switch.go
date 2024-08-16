package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda"
	default:
		return "erro"
	}
}

func main() {
	fmt.Println("Switchs")

	dia := diaDaSemana(0)
	fmt.Println(dia)
}
