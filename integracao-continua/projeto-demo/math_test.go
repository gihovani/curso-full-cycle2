package main

import "testing"

func TestSum(t *testing.T) {
	var result = 30
	var total = Sum(15, 15)
	if total != result {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, result)
	}
}

func TestSub(t *testing.T) {
	var result = 0
	var total = Sub(15, 15)
	if total != result {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado %d", total, result)
	}
}

func TestDiv(t *testing.T) {
	var result = 2
	var total = Div(6, 3)
	if total != result {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado %d", total, result)
	}
}

func TestTimes(t *testing.T) {
	var result = 6
	var total = Times(2, 3)
	if total != result {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado %d", total, result)
	}
}

func TestPow(t *testing.T) {
	var result = 256
	var total = Pow(2, 8)
	if total != result {
		t.Errorf("Resultado da Exponenciação é inválido: Resultado %d. Esperado %d", total, result)
	}
}
