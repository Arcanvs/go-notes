package main

import (
	"fmt"
	"strconv"
	"time"
)

func iterarSlice(data []int) {
	for index, value := range data {
		fmt.Printf("Este es el index: %d y este es el valor: %d\n", index, value)
	}
}

func main() {
	// Variables
	var x int
	x = 10
	y := 2
	fmt.Println(x)
	fmt.Println(y)
	// Captura de errores
	valorNum, errorMsg := strconv.ParseInt("3", 0, 64)
	if errorMsg != nil {
		fmt.Printf("%v\n", errorMsg)
	} else {
		fmt.Println(valorNum)
	}
	// Map
	dataMap := make(map[string]int)
	dataMap["key"] = 5
	fmt.Println(dataMap["key"])
	// Slice
	dataSlice := []int{1, 2, 3, 4, 5}
	// Print slice
	iterarSlice(dataSlice)
	// Append elemento al slice
	dataSlice = append(dataSlice, 6, 7, 8, 9, 10)
	// Print slice
	iterarSlice(dataSlice)
	// Crear Canal
	channel := make(chan int)
	// Llamar la rutina
	go doSomething(channel)
	<-channel
	// Punteros
	numero := 33
	fmt.Println(numero)
	numeroPuntero := &numero
	fmt.Println("Puntero espacio en memoria : ", numeroPuntero)

	fmt.Println("Puntero valor : ", *numeroPuntero)

}

// Go rutines
func doSomething(channel chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Listo")
	channel <- 1
}
