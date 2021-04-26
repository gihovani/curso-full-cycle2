package main

import "testing"

func TestSoma(t *testing.T) {
	var result = 30
	var total = Soma(15, 15)
	if total != result {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, result)
	}
}
