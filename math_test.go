package main
import "testing"

func SomaTest(t *testing.T) {
	totalSum := Soma(10, 10)

	if totalSum != 20 {
		t.Errorf(" Resultado %d. Esperado: %d", totalSum, 20)
	}
}