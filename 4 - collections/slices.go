package main

import "fmt"

func main() {
	//a declaracao do slice é igual a do array, apenas o numero que
	//determina o tamanho do array é omitido para não criar um array
	//ao invés de criar um slice, que não tem tamanho e pode crescer
	slice := []int {1, 10, 121}

	fmt.Println(slice)

	//adicionando mais items ao slice
	slice = append(slice, 4)

	fmt.Println(slice)
}
