package ejercicios

func RecursiveHeapSort[T any](arr []T, compare func(T, T) int) {
	size := len(arr)
	heapify(arr, size, compare)
	end := size - 1
	for end > 0 {
		arr[end], arr[0] = arr[0], arr[end]
		downHeap(arr, 0, end-1, compare)
		end--
	}
}

func heapify[T any](arr []T, size int, compare func(T, T) int) {
	start := (size - 2) / 2
	for start >= 0 {
		downHeap(arr, start, size - 1, compare)
		start--
	}
}

func downHeap[T any](arr []T, father, end int, compare func(T, T) int) {
	leftSon := father*2 + 1
	rightSon := leftSon + 1
	if leftSon <= end {
		if rightSon <= end && compare(arr[rightSon], arr[leftSon]) > 0 {
			leftSon = rightSon
		}
		if compare(arr[leftSon], arr[father]) > 0 {
			arr[leftSon], arr[father] = arr[father], arr[leftSon]
			downHeap(arr, leftSon, end, compare)
		}
	}
}