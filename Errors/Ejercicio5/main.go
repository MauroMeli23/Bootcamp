package main

import (
	"errors"
	"fmt"
)

func salaryCalculator(h int, s float64) (float64, error) {
	salary := (float64(h) * s) * 30

	if salary >= 150000 && h >= 80 {
		salary = salary * 0.2
		fmt.Println(salary)
		return salary, nil
	} else if h < 80 {
		err := errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		fmt.Println(err)
		return 0.0, err
	}
	fmt.Println(salary)
	return salary, nil
}

func main() {
	salary, err := salaryCalculator(85, 50.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salario:", salary)
	}
}
