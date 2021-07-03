package main

import "fmt"

func main() {
	//declarando uma constante
	const c = 3

	fmt.Println(c + 3)

	//usando com outros tipos de dados
	fmt.Println(3 + 1.2)

	//agora declarando um tipo para a constante
	const a int = 5

	//para usar em operações co outros numeros, tem que fazer cast
	fmt.Println(float32(a) + 4.1)
}