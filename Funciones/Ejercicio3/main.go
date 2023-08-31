package main

import "fmt"

func salaryCalculator(workeMinutes float64, category string) float64 {
	var salary float64
	workedHours := workeMinutes / 60
	if category == "C" {
		salary = workedHours * 1000.0
		fmt.Println(salary)
		return salary
	} else if category == "B" {
		salary = workedHours * 1500.0
		fmt.Println(salary + (salary * 0.2))
		return salary + (salary * 0.2)
	} else if category == "A" {
		salary = workedHours * 3000.0
		fmt.Println(salary + (salary * 0.5))
		return salary + (salary * 0.5)
	} else {
		fmt.Println("no ingreso nada")
	}
	return salary
}

func main() {
	salaryCalculator(60.0, "B")
}
