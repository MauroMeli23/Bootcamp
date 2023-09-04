package main

import "fmt"

func Increase(v *int) {
	//desreferenciamos el valor, entramos al casillero de la memoria y obtenemos lo que esta ahi
	*v++
}

func main() {
	var v = 19
	//Punteros: Un puntero es una variable que almacena la dirección de memoria de otra variable.
	//En Go, se declara un puntero utilizando el operador &, seguido de la variable a la que se desea apuntar
	var p *int = &v
	//tambien puede ser p := &v

	fmt.Printf("El puntero o referencia a la direccion de memeria: %v \n", p)
	fmt.Printf("Al desreferenciar el puntero p obtengo el valor: %d \n", *p)
	Increase(&v)
	fmt.Println("El valor de v ahora vale:", v)
}
