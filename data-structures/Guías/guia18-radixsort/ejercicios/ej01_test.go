package ejercicios

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEj1(test *testing.T) {
	casos := []struct {
		entrada    	[]int
		esperado	[]int
	}{
		{[]int{2, 1, 5, 7, 2, 5, 2, 8, 9, 0, 0, 1, 3, 5, 6}, []int{0, 0, 1, 1, 2, 2, 2, 3, 5, 5, 5, 6, 7, 8, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{0}, []int{0}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5}},
	}
	for _, caso := range casos {
		copiaEntrada := make([]int, len(caso.entrada))
		copy(copiaEntrada, caso.entrada)
		Ej1(copiaEntrada)
		assert.Equal(test, caso.esperado, copiaEntrada, "Para entrada %v", caso.entrada)
	}
}