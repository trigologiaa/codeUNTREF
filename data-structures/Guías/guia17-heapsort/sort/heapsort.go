package sort

// Función HeapSort implementa el algoritmo de ordenamiento HeapSort para ordenar un slice de cualquier tipo T, utilizando una función de comparación proporcionada.
func HeapSort[T any](arr []T, compare func(T, T) int) {
    size := len(arr)
    heapify(arr, compare)  // Llama a la función heapify para construir el heap inicial
    end := size - 1
    for end > 0 {
        arr[end], arr[0] = arr[0], arr[end]  // Intercambia el elemento raíz (máximo) con el último elemento del heap
        downHeap(arr, 0, end-1, compare)     // Reconstruye el heap después del intercambio
        end--
    }
}

// Función heapify construye un heap máximo a partir de un slice dado utilizando una función de comparación.
func heapify[T any](arr []T, compare func(T, T) int) {
    size := len(arr)
    start := (size - 2) / 2  // Índice del último nodo padre en el heap
    // Recorre todos los nodos padres del heap y realiza down-heap (reajuste descendente)
    for start >= 0 {
        downHeap(arr, start, size-1, compare)  // Llama a downHeap para asegurar que el subárbol en start sea un heap
        start--
    }
}

// Función downHeap ajusta el heap para mantener la propiedad del heap máximo desde el nodo padre dado hacia abajo.
func downHeap[T any](arr []T, start, end int, compare func(T, T) int) {
    father := start
    leftSon := father*2 + 1  // Índice del hijo izquierdo del padre
    rightSon := leftSon + 1  // Índice del hijo derecho del padre
    // Mientras el hijo izquierdo esté dentro del rango del heap
    for leftSon <= end {
        // Determina el hijo mayor (si existe)
        if rightSon <= end && compare(arr[rightSon], arr[leftSon]) > 0 {
            leftSon = rightSon  // Si el hijo derecho es mayor, se selecciona como el mayor
        }
        // Compara el hijo mayor con el padre y realiza el intercambio si el hijo es mayor
        if compare(arr[leftSon], arr[father]) > 0 {
            arr[leftSon], arr[father] = arr[father], arr[leftSon]
            father = leftSon  // Actualiza el índice del padre al hijo intercambiado
            leftSon = father*2 + 1  // Calcula el nuevo índice del hijo izquierdo
            rightSon = leftSon + 1  // Calcula el nuevo índice del hijo derecho
        } else {
            return  // Si el padre es mayor o igual que el hijo mayor, se mantiene la propiedad del heap
        }
    }
}
