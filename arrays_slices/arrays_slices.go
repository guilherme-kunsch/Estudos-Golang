package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//criação de array
	//inflexivel
	var array [5]int
	array[0] = 5
	array[1] = 54
	array[2] = 512
	array[3] = 5312
	array[4] = 53123
	fmt.Println(array)

	//inferência de tipos
	array2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array2)

	//criação do slice == uma fatia de um array
	slice := []int{10, 102, 20, 30}
	fmt.Println(slice)
	slice = append(slice, 10, 20, 40)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))

	//Arrays Internos
	//função make cria um array de 15 posição e vai retornar um slice das
	//10 primeiras posições desse array
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
