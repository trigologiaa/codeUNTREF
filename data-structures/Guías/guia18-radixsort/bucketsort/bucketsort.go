package main

// BucketSort ordena un arreglo de enteros utilizando el algoritmo
// de ordenamiento por urnas.
func BucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// Encontrar el valor maximo y minimo
	// para determinar el rango de los buckets
	max, min := arr[0], arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	// Crear los buckets
	numBuckets := (max - min) + 1
	buckets := make([][]int, numBuckets)

	// Colocar los elementos en los baldes correspondientes
	for _, num := range arr {
		buckets[num-min] = append(buckets[num-min], num)
	}

	// Combinar los buckets en el arreglo original
	var i int
	for _, bucket := range buckets {
		for _, num := range bucket {
			arr[i] = num
			i++
		}
	}
}
