package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Message string
}

//uso punteros, porque no quiero trabajar con una copia, lo que busco es que el valor de mi error sea el que le defino
//despues con &MyError

func (e *MyError) Error() string {
	return e.Message
}

func salaryCalculator(s int) error {
	if s < 150000 {
		return &MyError{
			Message: "Error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}
func main() {
	salary := 140000
	err := salaryCalculator(salary)

	if errors.Is(err, &MyError{}) {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
