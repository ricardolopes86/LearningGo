package main

import "fmt"

func main() {
	port := 3000
	//here its called arguments
	//we are using the write-only variable in here
	// if nothing is returned in the first parameter from the function,
	//we don't have to use it
	_, err := startWebserver(port, 2)
	fmt.Println(err)
}

//here its called parameters
//the second parenthesis are to determine what this function is going to return
//type of first return, error in case of anything goes wrong
func startWebserver(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting webserver")
	//codigo
	fmt.Println("Server started and its listening to port", port)
	fmt.Println("Number of retries in case of failure is", numberOfRetries)
	return port, nil
}
