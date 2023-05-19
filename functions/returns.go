package main

import "fmt"

// Funciones Variadica
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(values ...string) {
	for _, name := range values {
		fmt.Println(name)
	}
}

// Retornos con nombre
func getValues(x int) (doble int, triple int, cuadruple int) {
	doble = x * 2
	triple = x * 3
	cuadruple = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2))
	printNames("Cecilia", "Jorgito", "Amanda", "Gringo")
	fmt.Println(getValues(2))
}
