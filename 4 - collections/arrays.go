package main

import "fmt"

func main() {
	//arrays tem tamanho fixo, portanto o numero nos colchetes é o tamanho do array
	var arr [4]int
	//fazendo asignment de valores no array
	arr[0] = 1
	arr[1] = 12
	arr[2] = 123
	arr[3] = 1234

	//imprimindo o array inteiro
	fmt.Println(arr)

	//imprimindo apenas um índice do array
	fmt.Println(arr[1])

	//usando o formato curto
	arr2 := [4]int {2, 23, 234, 2345}
	fmt.Println(arr2)
	fmt.Println(arr2[3])
}
