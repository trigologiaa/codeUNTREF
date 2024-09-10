package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRadixSort(t *testing.T) {
	casosPrueba := []struct {
		entrada  	[]int
		esperado	[]int
	}{
		{[]int {170, 45, 75, 90, 802, 24, 2, 66}, []int {802, 170, 90, 75, 66, 45, 24, 2}},
		{[]int {5, 4, 3, 2, 1, 0}, []int {5, 4, 3, 2, 1, 0}},
		{[]int {1}, []int {1}},
		{[]int {}, []int {}},
		{[]int {9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int {9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, caso := range casosPrueba {
		entradaCopia := make([]int, len(caso.entrada))
		copy(entradaCopia, caso.entrada)
		RadixSort(entradaCopia)
		assert.Equal(t, caso.esperado, entradaCopia, "Para entrada %v", caso.entrada)
	}
}