package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("customers.txt")
	defer func() {
		fmt.Println("ejecución finalizada")
	}()
	if err != nil {
		panic(err.Error())
		fmt.Println("El archivo indicado no fue encontrado o está dañado")
	}

}
