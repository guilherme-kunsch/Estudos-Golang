package main

import (
	"fmt"
	"math"
)

//usamos interfaces quando precisamos ter flexibilidade com os tipos de dados

// criando método
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// criando interface
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

func main() {
	r := retangulo{10, 15}
	c := circulo{20}
	escreverArea(r)
	escreverArea(c)
}
