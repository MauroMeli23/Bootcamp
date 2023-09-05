package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("customers.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var content bytes.Buffer
	_, err = io.Copy(&content, file)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	fmt.Println(content.String())
}
