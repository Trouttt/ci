package main
import "testing"

func TestSoma(t *testing.T) {
	totalSum := Soma(15, 15)

	if totalSum != 30 {
		t.Errorf("Resultado %d. Esperado: %d", totalSum, 30)
	}
}