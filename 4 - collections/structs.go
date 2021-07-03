package main

import "fmt"

func main() {
	//definindo o struct e seus campos
	type user struct {
		id int
		nome string
		ultimoNome string
	}

	//instanciando um objeto como do tipo user, criado logo anteriormente
	var usuario user
	//atribuindo valores aos atributos do tipo user
	usuario.id = 1
	usuario.nome = "Ricardo"
	usuario.ultimoNome = "Silva"

	fmt.Println(usuario)

	//usando a forma contraída
	usuario1 := user{id: 2, nome: "Henrique", ultimoNome: "Garcia da Silva"}
	fmt.Println(usuario1)
}
