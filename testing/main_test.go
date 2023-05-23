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
