package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketSort(t *testing.T) {
	arr := []int{4, 12, 7, 5, 22, 9, 6, 3, 15, 3}
	expected := []int{3, 3, 4, 5, 6, 7, 9, 12, 15, 22}
	BucketSort(arr)
	fmt.Println(arr)
	assert.Equal(t, expected, arr)
}
