package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadixSort(t *testing.T) {
	arr := []int{1, 401, 11, 222, 257, 725, 214, 99, 100, 60, 315, 145, 273}
	expected := []int{1, 11, 60, 99, 100, 145, 214, 222, 257, 273, 315, 401, 725}
	RadixSort(arr)
	fmt.Println(arr)
	assert.Equal(t, expected, arr)
}
