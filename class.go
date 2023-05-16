package main

import "fmt"

// Struct === Class
type Empleado struct {
	id   int
	name string
}

// Receiver function
// Se asigna al Struct la funcion correspondiente
func (e *Empleado) SetId(id int) {
	e.id = id
}
func (e *Empleado) SetName(name string) {
	e.name = name
}
func (e *Empleado) GetId() int {
	return e.id
}
func (e *Empleado) GetName() string {
	return e.name
}

func main() {
	// Instancia Empleado 1
	empleado1 := Empleado{1, "Jorgito"}
	fmt.Println(empleado1)
	// Instancia Empleado 2
	empleado2 := Empleado{}
	empleado2.id = 2
	empleado2.name = "Andres"
	fmt.Println(empleado2)
	// Instancia Empleado 3
	empleado3 := Empleado{}
	empleado3.SetId(3)
	empleado3.SetName("Cecilia")
	fmt.Printf("Empleado id %v - nombre %v \n", empleado3.GetId(), empleado3.GetName())
}
