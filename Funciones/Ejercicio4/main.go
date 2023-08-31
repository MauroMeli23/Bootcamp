package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(opType string) (func(...int) (int, error), error) {
	switch opType {
	case minimum:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("No hay números para calcular el mínimo")
			}
			min := numbers[0]
			for _, num := range numbers {
				if num < min {
					min = num
				}
			}
			return min, nil
		}, nil

	case average:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("No hay números para calcular el promedio")
			}
			sum := 0
			for _, num := range numbers {
				sum += num
			}
			average := sum / len(numbers)
			return average, nil
		}, nil

	case maximum:
		return func(numbers ...int) (int, error) {
			if len(numbers) == 0 {
				return 0, errors.New("No hay números para calcular el máximo")
			}
			max := numbers[0]
			for _, num := range numbers {
				if num > max {
					max = num
				}
			}
			return max, nil
		}, nil

	default:
		return nil, errors.New("Operación no válida")
	}
}

func main() {
	minFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minValue, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue, _ := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Mínimo: %d\n", minValue)
	fmt.Printf("Promedio: %d\n", averageValue)
	fmt.Printf("Máximo: %d\n", maxValue)
}
