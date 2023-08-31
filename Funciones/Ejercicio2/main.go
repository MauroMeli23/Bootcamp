package main

import "fmt"

func promedy(marks []int) float64 {
	sum := 0
	for _, m := range marks {
		if m < 0 {
			panic("No se pueden introducir notas negativas.")
		}
		sum += m
	}
	promedio := float64(sum) / float64(len(marks))

	return promedio
}

func main() {
	marks := []int{5, 5, 10}
	promedio := promedy(marks)
	fmt.Printf("El promedio es %.2f\n", promedio)
}
