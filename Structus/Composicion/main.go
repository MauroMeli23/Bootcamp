package main

import "fmt"

type Autor struct {
	Nombre   string
	Apellido string
}

func (a Autor) nombreCompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo      string
	Descripcion string
	Autor       Autor
}

func (l Libro) informacion() {
	fmt.Println("Titulo:", l.Titulo)
	fmt.Println("Descripcion:", l.Descripcion)
	fmt.Println("Autor:", l.Autor.nombreCompleto())
}

func main() {
	autor := Autor{
		Nombre:   "Mauro",
		Apellido: "Iannizzotto",
	}

	libro := Libro{
		Titulo:      "Coding",
		Descripcion: "How to code",
		Autor:       autor,
	}
	libro.informacion()
}
