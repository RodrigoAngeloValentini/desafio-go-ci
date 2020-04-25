package main

import "testing"

func TestSoma(t *testing.T) {
	if Soma(5,5) != 10 {
		t.Error("Esperado 5+5 ser igual a 10")
	}
}