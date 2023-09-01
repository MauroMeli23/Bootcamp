package main

import "fmt"

func taxesCalculator(salary float64) float64 {
	var taxes float64

	if salary > 150000 {
		taxes = salary * 0.23
	} else if salary > 50000 {
		taxes = salary * 0.17
	}
	return taxes
}

func main() {
	salary := 60000.0
	taxes := taxesCalculator(salary)
	fmt.Println(taxes)
}
