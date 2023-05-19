package main

import "fmt"

type Persona struct {
	name string
	age  int
}

type Empleado struct {
	id int
}

// Composicion
type EmpleadoFullTime struct {
	Persona
	Empleado
}

func (empleadoFT EmpleadoFullTime) getMessage() string {
	return "Empleado a tiempo completo"
}

type EmpleadoPartTime struct {
	Persona
	Empleado
	taxRate int
}

func (empleadoPt EmpleadoPartTime) getMessage() string {
	return "Empleado a medio tiempo"
}

// Interface
type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	empleadoFt := EmpleadoFullTime{}
	empleadoFt.id = 1
	empleadoFt.name = "Jorgito"
	empleadoFt.age = 6
	fmt.Printf("%v\n", empleadoFt)

	empeladoPt := EmpleadoPartTime{}

	getMessage(empleadoFt)
	getMessage(empeladoPt)
}
