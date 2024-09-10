package ejercicios

import (
	"math/rand"
	"time"
)

func QuickSortPivoteAleatorio(array []int) []int {
	rand.Seed(time.Now().UnixNano())
	quickSortRandomPivot(array, 0, len(array)-1)
	return array
}

func quickSortRandomPivot(array []int, low, high int) {
	if low < high {
		pivotIndex := rand.Intn(high-low+1) + low
		pivotIndex = partitionRandomPivot(array, low, high, pivotIndex)
		quickSortRandomPivot(array, low, pivotIndex-1)
		quickSortRandomPivot(array, pivotIndex+1, high)
	}
}

func partitionRandomPivot(array []int, low, high, pivotIndex int) int {
	array[pivotIndex], array[high] = array[high], array[pivotIndex]
	pivot := array[high]
	i := low - 1
	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}