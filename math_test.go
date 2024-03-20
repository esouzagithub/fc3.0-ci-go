package main

import "testing"

func TestSomar(t *testing.T) {

	total := Somar(15, 10)

	if total != 30 {

		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
