package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArrayVacioQ(t *testing.T) {
	array := []int{}
	ordenado := []int{}
	assert.Equal(t, ordenado, QuickSortPivoteAleatorio(array))
}

func TestArrayUnElementoQ(t *testing.T) {
	array := []int{3}
	ordenado := []int{3}
	assert.Equal(t, ordenado, QuickSortPivoteAleatorio(array))
}

func TestCantidadParQ(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, ordenado, QuickSortPivoteAleatorio(array))
}

func TestCantidadImparQ(t *testing.T) {
	array := []int{5, 9, 2, 8, 1, 7, 4, 3, 6}
	ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, ordenado, QuickSortPivoteAleatorio(array))
}
