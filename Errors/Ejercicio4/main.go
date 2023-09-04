package main

import (
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func salaryCalculator(s int) error {
	if s < 150000 {
		return fmt.Errorf("Error: el salario ingresado no alcanza el mínimo imponible que es de %d", s)
	}
	return nil
}
func main() {
	salary := 140000
	err := salaryCalculator(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
