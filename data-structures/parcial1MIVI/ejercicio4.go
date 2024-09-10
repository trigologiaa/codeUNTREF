package parcial1MIVI

//Aplicar el teorema del maestro para calcular el orden de la búsqueda ternaria (similar a la búsqueda binaria, pero partiendo el arreglo en tres).

func TernarySearchRecursive(arr []int, left int, right int, key int) int {
	if left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3
		if arr[mid1] == key {
			return mid1
		}
		if arr[mid2] == key {
			return mid2
		}
		if key < arr[mid1] {
			return TernarySearchRecursive(arr, left, mid1-1, key)
		} else if key > arr[mid2] {
			return TernarySearchRecursive(arr, mid2+1, right, key)
		} else {
			return TernarySearchRecursive(arr, mid1+1, mid2-1, key)
		}
	}
	return -1
}