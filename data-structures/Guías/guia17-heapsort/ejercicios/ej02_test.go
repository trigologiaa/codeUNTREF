package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncontrarEnesimoMayor(t *testing.T) {
	array := []int {5, 9, 2, 8, 1, 7, 4, 3, 10, 6}
	expected := 8
	kesimo := EncontrarEnesimoMayor(array, 3)
	assert.Equal(t, expected, kesimo)
}