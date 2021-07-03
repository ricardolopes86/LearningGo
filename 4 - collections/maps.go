package main

import "fmt"

func main() {
	//declaracao simplificada de um map
	//entre colchetes é o tipo que as chaves terão
	//logo em seguida o tipo que os valores terão
	//e entre chaves, o map em si
	mapa := map[string]int{"chave1": 1}
	fmt.Println(mapa)

	fmt.Println(mapa["chave1"])

	//mudando o vlor de um item do map
	mapa["chave1"] = 32

	fmt.Println(mapa)

	//usando uma função builtin pra deletar um item do map
	delete(mapa, "chave1")

	fmt.Println(mapa)
}
