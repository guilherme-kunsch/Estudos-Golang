package main

import "fmt"

//metódos sempre vão estar ligados a alguma coisa

type usuario struct {
	nome  string
	idade int
}

//método
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade > 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"teste", 20}
	usuario1.salvar()

	usuario2 := usuario{"teste2", 22}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2)

}
