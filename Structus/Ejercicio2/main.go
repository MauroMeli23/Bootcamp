package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth int
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("ID: %d, Position: %s\n", e.ID, e.Position)
	fmt.Printf("Person - ID: %d, Name: %s, DateOfBirth: %d\n", e.Person.ID, e.Person.Name, e.Person.DateOfBirth)
}

func main() {
	firstPerson := Person{
		ID:          1,
		Name:        "Mauro",
		DateOfBirth: 11,
	}
	firstEmployee := Employee{
		ID:       1,
		Position: "Software Developer",
		Person:   firstPerson,
	}
	firstEmployee.PrintEmployee()
}
