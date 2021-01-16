package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground.")
	port := 8080
	_, err := startWebService(port, 2)
	fmt.Println(err)
}

func startWebService(port, numberOfRetrires int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port:", port)
	fmt.Println("Number of retries:", numberOfRetrires)
	return port, errors.New("Something went wrong")
}
