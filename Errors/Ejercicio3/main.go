package main

import (
	"errors"
	"fmt"
)

func salaryCalculator(s int) error {
	if s < 150000 {
		return errors.New("Error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}
func main() {
	salary := 1140000
	err := salaryCalculator(salary)

	if err != nil && err.Error() == "Error: el salario ingresado no alcanza el mínimo imponible" {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
