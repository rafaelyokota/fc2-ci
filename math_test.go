package main
import (
    "testing"
)

func TestSoma(t *testing.T){
	total := soma(15,15)
	esperado := 30
	if total != esperado{
		t.Errorf("Resultado soma invalido: Resultado %d", esperado)
	}
}
