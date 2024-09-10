package sort

// QuickSort ordena un array de elementos de cualquier tipo utilizando el algoritmo de quicksort.
// Recibe un array de elementos, una función de comparación y un entero que indica el orden de ordenamiento.
func QuickSort[T any](arr []T, compare func(a T, b T) int) {
	quickSort(arr, compare, 0, len(arr)-1) // Llama a la función privada quickSort para iniciar el ordenamiento desde el inicio (0) hasta el final (len(arr)-1).
}

// quickSort es la función privada que realiza la recursión para ordenar el array.
// Utiliza el algoritmo de quicksort dividiendo el array en subarrays y ordenándolos recursivamente.
func quickSort[T any](arr []T, compare func(a T, b T) int, low, high int) {
	if low < high { // Verifica si hay más de un elemento para ordenar en el subarray.
		pi := partition(arr, compare, low, high) // Llama a partition para obtener el índice del pivote.
		quickSort(arr, compare, low, pi-1)       // Ordena recursivamente el subarray izquierdo al pivote.
		quickSort(arr, compare, pi+1, high)      // Ordena recursivamente el subarray derecho al pivote.
	}
}

// partition realiza la partición del array alrededor del pivote y devuelve su índice.
// Utiliza la función de comparación para determinar la posición del pivote y la posición final del pivote en el array.
func partition[T any](arr []T, compare func(a T, b T) int, low, high int) int {
	pivot := arr[high] // Define el pivote como el último elemento del subarray.
	i := low - 1       // Inicializa el índice del elemento más pequeño.
	for j := low; j < high; j++ {
		if compare(arr[j], pivot) < 0 { // Compara cada elemento con el pivote utilizando la función de comparación.
			i++                         // Incrementa el índice del elemento más pequeño.
			arr[i], arr[j] = arr[j], arr[i] // Intercambia los elementos para colocar los menores al lado izquierdo del pivote.
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Coloca el pivote en su posición correcta, intercambiándolo con el primer elemento mayor.
	return i + 1                              // Retorna el índice del pivote.
}
