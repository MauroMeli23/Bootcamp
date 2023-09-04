package main

import (
	"errors"
	"fmt"
)

type Client struct {
	Legajo    int
	Nombre    string
	DNI       int
	Celular   int
	Domicilio string
}

var ExistingClients = []Client{
	{
		Legajo:    01,
		Nombre:    "Cliente1",
		DNI:       123,
		Celular:   321,
		Domicilio: "Calle 1",
	},
	{
		Legajo:    02,
		Nombre:    "Cliente2",
		DNI:       223,
		Celular:   421,
		Domicilio: "Calle 2",
	},
}

func AddClient(c Client) ([]Client, error) {
	for _, cl := range ExistingClients {
		if cl.DNI == c.DNI {
			panic("Error: el cliente ya existe")
		}
	}
	ExistingClients = append(ExistingClients, c)
	fmt.Println(ExistingClients)
	return ExistingClients, nil
}

func InputsCheck(c Client) (bool, error) {
	if c.Legajo == 0 || c.DNI == 0 || c.Nombre == "" || c.Celular == 0 || c.Domicilio == "" {
		return false, errors.New("Error: uno o más campos son cero")
	}
	return true, nil // Todos los campos son distintos de cero
}

func main() {
	defer func() {
		fmt.Println("Fin de la ejecución")
		if r := recover(); r != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
	}()

	newClient := Client{
		Legajo:    03,
		Nombre:    "Cliente 3",
		DNI:       323,
		Celular:   521,
		Domicilio: "Calle 3",
	}

	ok, err := InputsCheck(newClient)
	if err != nil {
		fmt.Println(err) // Manejar el error aquí
		return
	}

	if ok {
		_, err := AddClient(newClient)
		if err != nil {
			fmt.Println(err) // Manejar el error aquí
		}
	}
}
