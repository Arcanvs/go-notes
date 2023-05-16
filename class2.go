package main

import "fmt"

type Empleado struct {
	id       int
	name     string
	vacation bool
}

// Forma 4 "constructor"
func NewEmpleado(id int, name string, vacation bool) *Empleado {
	return &Empleado{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Forma 1 "constructor"
	empleado1 := Empleado{}
	fmt.Printf("%v \n", empleado1)
	// Forma 2 "constructor"
	empleado2 := Empleado{
		id:       1,
		name:     "Jorgito",
		vacation: true,
	}
	fmt.Printf("%v\n", empleado2)
	// Forma 3 "constructor"
	empleado3 := new(Empleado)
	fmt.Printf("%v\n", *empleado3)
	empleado3.id = 2
	empleado3.name = "Andres"
	fmt.Printf("%v\n", *empleado3)
	// Forma 4 "constructor"
	empleado4 := NewEmpleado(3, "Cecilia", true)
	fmt.Printf("%v\n", *empleado4)
}
