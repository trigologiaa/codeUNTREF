package ejercicios

import(
	"ayp2/ordava/sort"
)

func EncontrarEnesimoMayor(array []int, n int) int {
	compare := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	}
	sort.HeapSort(array, compare)
	return array[n - 1]
}