package ejercicios_avl

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortingOnEmptyArray(t *testing.T) {
	elements := []int{}
	sorted := AVLTreeSort(elements)
	assert.Equal(t, []int{}, sorted)
}

func TestSortingOnSingleElementArray(t *testing.T) {
	elements := []int{1}
	sorted := AVLTreeSort(elements)
	assert.Equal(t, []int{1}, sorted)
}

func TestSortingOnMultipleElementArray(t *testing.T) {
	elements := []int{4, 2, 6, 1, 3, 5, 7}
	sorted := AVLTreeSort(elements)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, sorted)
}