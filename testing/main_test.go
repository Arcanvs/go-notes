package main

import "testing"

func TestSum(t *testing.T) {
	/* total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta obtiene %d y se espera %d", total, 10)
	} */
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{10, 1, 11},
	}

	for _, item := range tables {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta obtiene %d y se espera %d", total, item.c)
		}
	}
}

func TestMayor(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{6, 2, 6},
		{6, 9, 9},
	}

	for _, item := range tables {
		mayor := GetMayor(item.a, item.b)
		if mayor != item.c {
			t.Errorf("GetMayor incorrecto obtiene %d y se espera %d", mayor, item.c)
		}
	}

}

/*
	A tener en cuenta
	go test -coverprofile=coverage.out
	permite conocer que parte del codigo esta testeado

	Para leer las metricas de forma legible
	go tool cover -func=coverage.out
	Resumen en la terminal
	go tool cover -html=coverage.out
	Resumen en el navegador
*/
