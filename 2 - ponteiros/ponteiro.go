package main

import "fmt"

func main() {
	//declarando um ponteiro
	var nome *string = new(string)

	//a forma correta de fazer um assignment
	*nome = "Ricardo Silva"

	//usando o ponteiro para pegar o valor do endereco de memória
	fmt.Println("O nome eh ", *nome)

	//usando o ponteiro para pegar o ENDEREÇO de memória
	fmt.Println(nome)

	//agora usando um derefference, onde vamos criar o ponteiro depois e apontar para uma variavel já existente
	primeiroNome := "Ricardo"
	fmt.Println(primeiroNome)

	//criamos o ponteiro já apontando para a variavel que já existe
	ptr := &primeiroNome
	fmt.Println(primeiroNome, ptr)

	//mudamos o valor da variavel
	primeiroNome = "Primeiro NOme"

	//imprimindo o valor
	fmt.Println(primeiroNome, ptr)
}
