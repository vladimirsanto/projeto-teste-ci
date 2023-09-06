package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 14)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Rsultado: %d. Esperado: %d", total, 30)
	}
}
